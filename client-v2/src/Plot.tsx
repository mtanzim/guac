import { useEffect, useState } from "react";
import { DailyChart } from "./DailyChart";
import { StatsData } from "./data-types";
import { LanguageChart } from "./LanguageChart";
import { ProjectChart } from "./ProjectChart";
import { BASE_URL, DEFAULT_DAY_RANGE, getDateRange } from "./utils";
import { LanguagePct } from "./LanguagePct";

export function Plot({
  onLogout,
  token,
}: {
  onLogout: () => void;
  token: string;
}) {
  const [data, setData] = useState<null | StatsData>(null);
  const [loading, setLoading] = useState(false);
  const [errMsg, setErrMsg] = useState<null | string>(null);
  useEffect(() => {
    const { starting, ending } = getDateRange(DEFAULT_DAY_RANGE);
    const url =
      starting && ending
        ? `${BASE_URL}/api/v1/data?start=${starting}&end=${ending}`
        : `${BASE_URL}/api/v1/data`;
    setLoading(true);
    fetch(url, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((res) => {
        if (res.ok) {
          return res.json();
        }
        throw new Error("cannot fetch");
      })
      .then((d) => setData(d))
      .catch((err) => {
        console.log(err);
        setErrMsg("Something went wrong");
        onLogout();
      })
      .finally(() => setLoading(false));
  }, [token, onLogout]);

  if (loading) {
    return <p>Loading...</p>;
  }
  if (errMsg) {
    return <p>{errMsg}</p>;
  }
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4">
      {/* <code>{JSON.stringify(data?.dailyDuration, null, 2)}</code> */}
      {data?.dailyDuration && (
        <DailyChart dailyDuration={data?.dailyDuration} />
      )}
      {data?.languageStats && (
        <LanguageChart languageDurations={data?.languageStats?.durations} />
      )}
      {data?.projectStats && (
        <ProjectChart projectDurations={data?.projectStats?.durations} />
      )}
      {data?.languageStats?.percentages && (
        <LanguagePct percentages={data?.languageStats?.percentages} />
      )}
    </div>
  );
}
