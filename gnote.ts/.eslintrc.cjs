module.exports = {
  extends: ['standard-with-typescript', 'prettier'],
  parserOptions: { project: './tsconfig.json' },
  rules: {
    '@typescript-eslint/no-misused-promises': ['off'],
  },
}
