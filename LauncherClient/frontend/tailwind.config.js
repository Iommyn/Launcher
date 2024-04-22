/** @type {import('tailwindcss').Config} */
const plugin = require("tailwindcss/plugin");

export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],


  theme: {
    screens: {
      auth_width_launcher: {'max': "522px"},
    },

    extend: {
      backgroundImage: {
        'launcher': 'linear-gradient(39.28deg, #1B2130 43.01%, #122F4F 65.35%, #18575B 91.87%), linear-gradient(180deg, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.27) 100%)',
        'sidebar': 'linear-gradient(180deg, #182A46 0%, #182F54 100%)',
        'img-panel': "url('/src/assets/images/bg_panel.png')",
        'page_info': 'linear-gradient(39.28deg, #1B2130 43.01%, #122F4F 65.35%, #18235B 91.87%), linear-gradient(180deg, rgba(0, 0, 0, 0) 0%, rgba(0, 0, 0, 0.27) 100%)',
      },
      colors: {
        "btn_auth": "rgba(57, 88, 254, 1)",
        "link": "#57A4FF",
        "panel_bg": "rgb(30,49,81)",
      },
      fontFamily: {
        RubikSemiBold: ["RubikSemiBold"],
        RubikRegular: ["RubikRegular"],
      },
    },
  },
  plugins: [
    plugin(function ({ addBase }) {
      addBase({
        "@font-face": {
          fontFamily: "RubikSemiBold",
          src: 'url(/fonts/Rubik-SemiBold.woff) format("woff")',
        },
      });
    }),
    plugin(function ({ addBase }) {
      addBase({
        "@font-face": {
          fontFamily: "RubikRegular",
          src: 'url(/fonts/Rubik-Regular.woff) format("woff")',
        },
      });
    }),
  ],
}