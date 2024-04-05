module.exports = {
  "root": true,
  "plugins": [
    import("prettier-plugin-tailwindcss"),
    import("@ianvs/prettier-plugin-sort-imports")
  ],
  "printWidth": 80,
  "tabWidth": 2,
  "singleQuote": false,
  "trailingComma": "all",
  "arrowParens": "always",
  "semi": false,
  "endOfLine": "auto",
  "tailwindFunctions": ["clsx"],
}