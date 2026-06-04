/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./template/*.{html,js,templ}', './post/styles.html'],
  theme: {
    extend: {
      fontFamily: {
        // new web fonts first, the previous defaults kept as fallbacks
        serif: [
          '"Source Serif 4"',
          'ui-serif',
          'Georgia',
          'Cambria',
          '"Times New Roman"',
          'Times',
          'serif',
        ],
        mono: [
          '"JetBrains Mono"',
          'ui-monospace',
          'SFMono-Regular',
          'Menlo',
          'Monaco',
          'Consolas',
          '"Liberation Mono"',
          '"Courier New"',
          'monospace',
        ],
      },
    },
  },
  plugins: [],
};
