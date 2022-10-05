/** @type {import("tailwindcss").Config} */

module.exports = {
  content: ["./app/ui/views/*.html", "./app/ui/views/partials/*.html"],
  daisyui: {
    themes: ["bumblebee"],
  },
  plugins: [require("daisyui"), require('@tailwindcss/forms')],
}
