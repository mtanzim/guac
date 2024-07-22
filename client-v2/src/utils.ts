import * as colors from "./colors.json";

export const toCustomDateStr = (d: string): string => {
  return new Date(d).toLocaleDateString(undefined, {
    month: "short",
    day: "2-digit",
  });
};

export function getDateRange(days: number) {
  const formatDateForReq = (date: Date) => {
    return date.toISOString().split("T")[0];
  };

  const endDate = new Date();
  endDate.setDate(endDate.getDate() - 1);
  const ending = formatDateForReq(endDate);
  const startDate = new Date();
  startDate.setDate(startDate.getDate() - days);
  const starting = formatDateForReq(startDate);
  return { starting, ending };
}

export function getColor(lang: string): string {
  return (
    (colors as Record<string, { color: string | null }>)?.[lang]?.color ||
    "gray"
  );
}

export const DEFAULT_DAY_RANGE = 7;
export const TOKEN_KEY = "WakaToken";
export const BASE_URL = "http://localhost:8080";
export const TOP_N_LANGUAGES = 7;
export const TOP_N_PROJECTS = 7;
