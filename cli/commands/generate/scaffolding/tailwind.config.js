/** @type {import('tailwindcss').Config} */


module.exports = {
    content: [
      // "../**/*.{html,templ,go,js}",
      "./views/**/*/*.{templ,html,js}",
    ],
    theme: {
      extend: {},
    },
    plugins: [
      require('@tailwindcss/forms'),
    ],
  }
  
  