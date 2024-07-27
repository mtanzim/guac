import { useEffect, useState } from "react";
import { DailyChart } from "./DailyChart";
import { StatsData } from "./data-types";
import { LanguageChart } from "./LanguageChart";
import { LanguagePct } from "./LanguagePct";
import { ProjectChart } from "./ProjectChart";
import { BASE_URL, formatDateForReq } from "./utils";

export function Plot({
  onLogout,
  token,
  start,
  end,
}: {
  onLogout: () => void;
  token: string;
  start: Date;
  end: Date;
}) {
  const [data, setData] = useState<null | StatsData>(null);
  const [loading, setLoading] = useState(false);
  const [errMsg, setErrMsg] = useState<null | string>(null);

  useEffect(() => {
    const starting = formatDateForReq(start);
    const ending = formatDateForReq(end);
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
  }, [token, onLogout, start, end]);

  if (errMsg) {
    return <p>{errMsg}</p>;
  }
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4 p-4">
      <DailyChart loading={loading} dailyDuration={data?.dailyDuration} />
      <ProjectChart
        loading={loading}
        projectDurations={data?.projectStats?.durations}
      />
      <LanguageChart
        loading={loading}
        languageDurations={data?.languageStats?.durations}
      />
      <LanguagePct
        loading={loading}
        rawPercentages={data?.languageStats?.percentages}
      />
    </div>
  );
}
