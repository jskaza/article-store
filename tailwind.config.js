/** @type {import("tailwindcss").Config} */

module.exports = {
  content: ["./app/ui/views/*.html", "./app/ui/views/partials/*.html"],
  daisyui: {
    themes: ["autumn"],
  },
  plugins: [require("daisyui"),  require('@tailwindcss/forms')],
}
