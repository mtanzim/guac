export const toCustomDateStr = (d: string): string => {
  return new Date(d).toLocaleDateString(undefined, {
    month: "short",
    day: "2-digit",
  });
};