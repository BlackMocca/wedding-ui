/** @type {import('tailwindcss').Config} */
var defaultColorConfig = require('./tailwind.config-color.js')

module.exports = {
  mode: 'aot',
  content: [
    "../domain/**.go",
    "../domain/**/*.go",
    "../pages/**.go",
    "../pages/**/*.go",
  ],
  theme: {
    screens: {
      iphone: { 
        min: "280px", max: "639px" 
      },
      tablet: "640px",
      laptop: "1280px",
    },
    fontFamily: {
      regular: "Prompt-Regular",
      medium: "Prompt-Medium",
    },
    fontSize: {
      sm: '0.8rem',
      base: '1rem',
      xl: '1.25rem',
      '2xl': '1.563rem',
      '3xl': '1.953rem',
      '4xl': '2.441rem',
      '5xl': '3.052rem',
    },
    colors: {
      primary: {
        base: "#494342"
      },
      secondary: {
        base: "#FFFFFF"
      },
      red: "red",
      chacoal: "#333333"
    },
    extend: {
      fontFamily: {
        sans: ['Prompt-Regular', 'Prompt-Medium']
      },
      borderWidth: {
        '0.5': '0.5px',
      },
      colors:{
        ...defaultColorConfig.theme.colors
      },
    }
  },
  plugins: [],
  safelist: [],
}

