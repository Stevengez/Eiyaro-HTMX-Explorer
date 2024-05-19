/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./public/**/*.html",
    "./templates/**/*.html"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require("@tailwindcss/typography"),
    require('daisyui'),
  ],
  daisyui: {
    themes: [
      {
        eiyaro: {
          "primary": "#1977d6",
          "secondary": "#0284c7",
          "accent": "#7dd3fc",
          "neutral": "#fff",
          "neutral-content": "#20252d",
          "base-100": "#f2f2f2",
          "base-300": "#262626",
          "info": "#22d3ee",
          "success": "#10b981",
          "warning": "#f59e0b",
          "error": "#f43f5e"
        },
      },
      "dark"
    ]
  }
}

