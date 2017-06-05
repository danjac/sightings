export const removeTrailingComma = (str) => {
  const trimmed = str.trim();
  return trimmed.endsWith(",") ? trimmed.slice(0, -1) : trimmed;
}
