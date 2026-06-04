(function () {
  // getting the canvas in the title and closing if it did not render
  const canvas = document.getElementById("word-cycle");
  if (!canvas || !canvas.getContext) return;

  // single source of truth: the words come from the canvas's data-words attr
  const WORDS = (canvas.dataset.words || "")
    .split(",")
    .map(function (w) {
      return w.trim();
    })
    .filter(Boolean);
  if (!WORDS.length) return;

  // derive the accessible name from the same list
  canvas.setAttribute("role", "img");
  canvas.setAttribute("aria-label", WORDS.join(", "));

  const COLOR = "#34d399"; // emerald-400, matches the surrounding text
  const FONT = 'italic %FONTPX%px Georgia, Cambria, "Times New Roman", serif';
  const HOLD = 2500; // ms a word stays crisp
  const MORPH = 1500; // ms for the full pixelate -> reform transition
  const RAMP = 0.3; // each char's ramp length, as a fraction of the morph (0..0.5)
  // window (in morph progress) over which the canvas width animates old -> new
  const RESIZE_LO = 0.45;
  const RESIZE_HI = 0.55;

  const ctx = canvas.getContext("2d");
  const buf = document.createElement("canvas"); // full-res glyph render
  const bctx = buf.getContext("2d");
  const tmp = document.createElement("canvas"); // tiny downsampled buffer
  const tctx = tmp.getContext("2d");

  let idx = 0;
  let fontPx = 48;
  let dpr = 1;
  let pad = 8; // horizontal breathing room (covers italic overhang)
  let widthCss = 0; // current canvas width in css px (animated)
  let heightCss = 0;
  let baselineY = 0; // y of the alphabetic baseline inside the canvas
  let maxBlock = 10; // largest pixel-block size, scaled to the font
  let metrics = {}; // word -> { adv: [glyph widths], sum, w (= sum + pad) }

  // transition state
  let morphing = false;
  let morphStart = 0;
  let phaseStart = 0;
  let oldWord = "";
  let newWord = "";
  let outStart = []; // per-char dissolve-out start time (0..0.5)
  let inStart = []; // per-char reform start time (0.5..1)
  let outBlock = []; // per-char chunkiness while dissolving out
  let inBlock = []; // per-char chunkiness while reforming

  function font() {
    return FONT.replace("%FONTPX%", fontPx);
  }

  function clamp01(v) {
    return v < 0 ? 0 : v > 1 ? 1 : v;
  }

  function computeMetrics() {
    bctx.setTransform(1, 0, 0, 1, 0, 0);
    bctx.font = font();
    pad = Math.round(fontPx * 0.18);
    metrics = {};
    for (let i = 0; i < WORDS.length; i++) {
      const word = WORDS[i];
      const adv = [];
      let sum = 0;
      for (let j = 0; j < word.length; j++) {
        const w = bctx.measureText(word[j]).width;
        adv.push(w);
        sum += w;
      }
      metrics[word] = { adv: adv, sum: sum, w: Math.ceil(sum) + pad };
    }
  }

  function measure() {
    const cs = window.getComputedStyle(canvas);
    fontPx = parseFloat(cs.fontSize) || 48;
    dpr = Math.max(1, window.devicePixelRatio || 1);
    maxBlock = Math.max(6, Math.round(fontPx / 6));

    // derive the real font metrics so the word sits on the text baseline
    bctx.setTransform(1, 0, 0, 1, 0, 0);
    bctx.font = font();
    const tm = bctx.measureText("Mpqgy");
    const asc = tm.fontBoundingBoxAscent || fontPx * 0.8;
    const desc = tm.fontBoundingBoxDescent || fontPx * 0.2;
    const vpad = Math.round(fontPx * 0.06); // tiny top/bottom safety margin
    heightCss = Math.ceil(asc + desc) + vpad * 2;
    baselineY = vpad + asc;
    // shift the inline canvas down so its drawn baseline meets the line's
    // baseline (canvas's own inline baseline is its bottom edge)
    canvas.style.verticalAlign = -(desc + vpad) + "px";

    computeMetrics();
    widthCss = -1; // force a resize on the next frame
  }

  function setCanvasWidth(wCss) {
    wCss = Math.round(wCss);
    if (wCss === widthCss) return;
    widthCss = wCss;
    canvas.style.width = wCss + "px";
    canvas.style.height = heightCss + "px";
    canvas.width = buf.width = Math.round(wCss * dpr);
    canvas.height = buf.height = Math.round(heightCss * dpr);
  }

  // left x-position of each character so the word is centred in the canvas
  function layout(word) {
    const m = metrics[word];
    let x = (widthCss - m.sum) / 2;
    const xs = [];
    for (let i = 0; i < word.length; i++) {
      xs.push(x);
      x += m.adv[i];
    }
    return xs;
  }

  function drawGlyph(ch, x) {
    bctx.setTransform(1, 0, 0, 1, 0, 0);
    bctx.clearRect(0, 0, buf.width, buf.height);
    bctx.scale(dpr, dpr);
    bctx.font = font();
    bctx.fillStyle = COLOR;
    bctx.textAlign = "left";
    bctx.textBaseline = "alphabetic";
    bctx.fillText(ch, x, baselineY);
  }

  // composite the current `buf` onto ctx, pixelated by `block` (1 = crisp)
  function blit(block) {
    if (block <= 1.01) {
      ctx.imageSmoothingEnabled = false;
      ctx.drawImage(buf, 0, 0);
      return;
    }
    const sw = Math.max(1, Math.round(widthCss / block));
    const sh = Math.max(1, Math.round(heightCss / block));
    tmp.width = sw;
    tmp.height = sh;
    tctx.imageSmoothingEnabled = true; // averaging down-sample
    tctx.clearRect(0, 0, sw, sh);
    tctx.drawImage(buf, 0, 0, buf.width, buf.height, 0, 0, sw, sh);
    ctx.imageSmoothingEnabled = false; // nearest-neighbour up-scale = blocks
    ctx.drawImage(tmp, 0, 0, sw, sh, 0, 0, canvas.width, canvas.height);
  }

  function renderCrisp(word) {
    const xs = layout(word);
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    for (let i = 0; i < word.length; i++) {
      drawGlyph(word[i], xs[i]);
      blit(1);
    }
  }

  // render `word` where each character has its own pixelation amount (0..1)
  function renderStaggered(word, pix, blocks) {
    const xs = layout(word);
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    for (let i = 0; i < word.length; i++) {
      drawGlyph(word[i], xs[i]);
      blit(1 + pix[i] * (blocks[i] - 1));
    }
  }

  function randomStarts(len, lo, hi) {
    const arr = [];
    for (let i = 0; i < len; i++) arr.push(lo + Math.random() * (hi - lo));
    return arr;
  }

  function randomBlocks(len) {
    const arr = [];
    for (let i = 0; i < len; i++)
      arr.push(maxBlock * (0.7 + Math.random() * 0.6));
    return arr;
  }

  function beginMorph(now) {
    morphing = true;
    morphStart = now;
    oldWord = WORDS[idx];
    newWord = WORDS[(idx + 1) % WORDS.length];
    // dissolve-out finishes by the midpoint; reform starts at the midpoint
    outStart = randomStarts(oldWord.length, 0, 0.5 - RAMP);
    inStart = randomStarts(newWord.length, 0.5, 1 - RAMP);
    outBlock = randomBlocks(oldWord.length);
    inBlock = randomBlocks(newWord.length);
  }

  // canvas width for the current moment: tight per word, eased old->new at the
  // (fully pixelated) midpoint so the surrounding text reflows smoothly
  function displayWidth(now) {
    if (!morphing) return metrics[WORDS[idx]].w;
    const p = (now - morphStart) / MORPH;
    const ow = metrics[oldWord].w;
    const nw = metrics[newWord].w;
    if (p <= RESIZE_LO) return ow;
    if (p >= RESIZE_HI) return nw;
    let t = (p - RESIZE_LO) / (RESIZE_HI - RESIZE_LO);
    t = t * t * (3 - 2 * t); // smoothstep
    return ow + (nw - ow) * t;
  }

  function frame(now) {
    if (!phaseStart) phaseStart = now;
    if (!morphing && now - phaseStart >= HOLD) beginMorph(now);

    let word = WORDS[idx];
    let pix = null;
    let blocks = null;

    if (morphing) {
      const p = (now - morphStart) / MORPH;
      if (p >= 1) {
        idx = (idx + 1) % WORDS.length;
        morphing = false;
        phaseStart = now;
        word = WORDS[idx];
      } else if (p < 0.5) {
        word = oldWord;
        blocks = outBlock;
        pix = [];
        for (let i = 0; i < oldWord.length; i++) {
          pix.push(clamp01((p - outStart[i]) / RAMP)); // crisp -> blocks
        }
      } else {
        word = newWord;
        blocks = inBlock;
        pix = [];
        for (let j = 0; j < newWord.length; j++) {
          pix.push(1 - clamp01((p - inStart[j]) / RAMP)); // blocks -> crisp
        }
      }
    }

    setCanvasWidth(displayWidth(now));
    if (pix) renderStaggered(word, pix, blocks);
    else renderCrisp(word);

    requestAnimationFrame(frame);
  }

  function start() {
    measure();
    const reduce =
      window.matchMedia &&
      window.matchMedia("(prefers-reduced-motion: reduce)").matches;
    if (reduce) {
      setCanvasWidth(metrics[WORDS[0]].w);
      renderCrisp(WORDS[0]); // honour reduced-motion: just show the first word
      return;
    }
    requestAnimationFrame(frame);
  }

  let resizeTimer;
  window.addEventListener("resize", function () {
    clearTimeout(resizeTimer);
    resizeTimer = setTimeout(measure, 150);
  });

  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", start);
  } else {
    start();
  }
})();
