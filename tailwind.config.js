/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['views/*.html',],
  theme: {
    extend: {
      colors: {
        'cream': '#EFDECD',
        'lime': '#8FBC8F',
        'a-bottom': '#E3E4DF',
        'a-top': '#D0C6BB',
      },
      fontFamily: {
        pragmatica: ['Pragmatica'],
        familjen: ['Familjen Grotesk'],
      },
    },
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["light", "dark", "luxury", "sunset"],
  },
}
    
