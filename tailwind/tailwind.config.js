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
    fontSize: {
      sm: '0.875rem',
      base: '1.125rem',
      xl: '1.25rem',
      '2xl': '1.375rem',
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
        sans: ['Prompt-Regular', 'Prompt-Medium', "sans-serif"],
        regular: 'Prompt-Regular',
        medium: 'Prompt-Medium',
      },
      colors:{
        ...defaultColorConfig.theme.colors
      },
    }
  },
  plugins: [],
  safelist: [],
}

