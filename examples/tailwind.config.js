/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "../components/**/*/*/{templ,html,js}",
    "./views/**/*/*.{templ,html,js}",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
  ],
}

