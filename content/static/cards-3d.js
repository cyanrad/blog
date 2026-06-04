// Cursor-following 3D tilt for blog cards.
// .card-3d: Is the perspective class (parent), used to ease the inner tilting card for smooth hovering
// .card-tilt: The actual tilting class that rotates the element
(function () {
  // cases to disable the rotation
  if (
    window.matchMedia && // checking if function exists for very old browsers
    (window.matchMedia("(hover: none)").matches || // skipping tilt if the device can't (like a touch screen)
      window.matchMedia("(prefers-reduced-motion: reduce)").matches) // when the user has "reduced motion" setting in their OS
  ) {
    return;
  }

  const MAX_TILT = 6; // max tilt in degrees
  const EASE = 0.2; // how much of the remaining distance to close each frame
  const cards = document.querySelectorAll(".card-3d");

  for (let i = 0; i < cards.length; i++) {
    bind(cards[i]);
  }

  function bind(card) {
    const tilt = card.querySelector(".card-tilt"); // getting the first card tilt element inside the card
    if (!tilt) return;

    // currently applied rotation (deg)
    let curX = 0;
    let curY = 0;
    // target rotation driven by the cursor (deg)
    let tgtX = 0;
    let tgtY = 0;

    let active = false; // if pointer is over the card
    let raf = null; // request animation frame

    card.addEventListener("mouseenter", function () {
      active = true;
      ensureLoop();
    });

    card.addEventListener("mousemove", function (e) {
      // getting card boundry
      const r = card.getBoundingClientRect();

      // Getting mouse position relative to the center of the card in percentages
      const px = ((e.clientX - r.left) / r.width - 0.5) * 2; // -1 .. 1
      const py = ((e.clientY - r.top) / r.height - 0.5) * 2; // -1 .. 1

      tgtY = px * MAX_TILT;
      tgtX = -py * MAX_TILT;
      ensureLoop();
    });

    card.addEventListener("mouseleave", function () {
      active = false;
      tgtX = 0;
      tgtY = 0;
      ensureLoop(); // still loop to ease back to flat
    });

    // prevents recursive functions repeating
    function ensureLoop() {
      if (raf === null) raf = requestAnimationFrame(loop);
    }

    // the rotation animation logic
    function loop() {
      // setting the new current rotation
      curX += (tgtX - curX) * EASE;
      curY += (tgtY - curY) * EASE;

      // Pause rotation once we're close enough
      const done = Math.abs(tgtX - curX) < 0.01 && Math.abs(tgtY - curY) < 0.01;
      if (done) {
        curX = tgtX;
        curY = tgtY;
      }

      tilt.style.transform =
        "rotateY(" + curY + "deg) rotateX(" + curX + "deg)";

      // keep animating while hovered, or until the ease-back finishes
      if (active || !done) {
        raf = requestAnimationFrame(loop);
      } else {
        raf = null;
      }
    }
  }
})();
