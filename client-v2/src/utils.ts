import * as colors from "./colors.json";
import { format } from "date-fns";

const fixDateOffset = (d: string): Date => {
  const date = new Date(d);
  const userTimezoneOffset = date.getTimezoneOffset() * 60000;
  return new Date(date.getTime() + userTimezoneOffset);
};

export const toCustomDateStr = (d: string): string => {
  return format(fixDateOffset(d), "MMM dd");
};

export const toCustomDateStrWithYear = (d: string): string => {
  return format(fixDateOffset(d), "MMM dd, yyyy");
};

export const formatDateForReq = (date: Date) => {
  return date.toISOString().split("T")[0];
};

export function getColor(lang: string): string {
  return (
    (colors as Record<string, { color: string | null }>)?.[lang]?.color ||
    "gray"
  );
}

export const DEFAULT_DAY_RANGE = 7;
export const TOKEN_KEY = "WakaToken";
export const BASE_URL = import.meta.env.PROD ? "" : "http://localhost:8080";
export const TOP_N_LANGUAGES = 10;
export const TOP_N_PROJECTS = 7;
