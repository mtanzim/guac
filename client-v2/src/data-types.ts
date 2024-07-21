type DailyDurationItem = {
  date: string;
  minutes: number;
};

type LanguageDuration = {
  language: string;
  minutes: number;
};

type LanguagePercentage = {
  language: string;
  percentage: number;
};

type ProjectDuration = {
  project: string;
  minutes: number;
};

export type StatsData = {
  dailyDuration: DailyDurationItem[];
  languageStats: {
    durations: LanguageDuration[];
    percentages: LanguagePercentage[];
  };
  projectStats: {
    durations: ProjectDuration[];
  };
  startDate: string;
  endDate: string;
};

export const example: StatsData = {
  dailyDuration: [
    {
      date: "2024-07-14",
      minutes: 0,
    },
    {
      date: "2024-07-15",
      minutes: 237.99552463333333,
    },
    {
      date: "2024-07-16",
      minutes: 167.64883861666667,
    },
    {
      date: "2024-07-17",
      minutes: 197.18443561666666,
    },
    {
      date: "2024-07-18",
      minutes: 252.80808461666666,
    },
    {
      date: "2024-07-19",
      minutes: 276.7612009,
    },
    {
      date: "2024-07-20",
      minutes: 58.788562166666665,
    },
  ],
  languageStats: {
    durations: [
      {
        language: "Python",
        minutes: 595.6876050833333,
      },
      {
        language: "Go",
        minutes: 284.69743489999996,
      },
      {
        language: "TypeScript",
        minutes: 132.42725261666666,
      },
      {
        language: "JSON",
        minutes: 44.86217698333333,
      },
      {
        language: "JavaScript",
        minutes: 34.23529576666667,
      },
      {
        language: "Other",
        minutes: 29.891053916666664,
      },
      {
        language: "YAML",
        minutes: 22.008704066666667,
      },
      {
        language: "Markdown",
        minutes: 15.427683333333334,
      },
      {
        language: "Docker",
        minutes: 12.029500500000001,
      },
      {
        language: "Makefile",
        minutes: 11.689532849999999,
      },
      {
        language: "Bash",
        minutes: 2.7284651333333327,
      },
      {
        language: "HTML",
        minutes: 2.051391466666667,
      },
      {
        language: "INI",
        minutes: 1.1178621833333335,
      },
      {
        language: "TOML",
        minutes: 1.0208547333333333,
      },
      {
        language: "Prisma",
        minutes: 0.49474311666666665,
      },
      {
        language: "Git",
        minutes: 0.3268461833333333,
      },
      {
        language: "Git Config",
        minutes: 0.20679083333333334,
      },
      {
        language: "TSConfig",
        minutes: 0.1655883833333333,
      },
      {
        language: "CSS",
        minutes: 0.1178645,
      },
    ],
    percentages: [
      {
        language: "Python",
        percentage: 50.00791494839255,
      },
      {
        language: "Go",
        percentage: 23.900321223761285,
      },
      {
        language: "TypeScript",
        percentage: 11.117254630096127,
      },
      {
        language: "JSON",
        percentage: 3.766175276835614,
      },
      {
        language: "JavaScript",
        percentage: 2.8740496601285264,
      },
      {
        language: "Other",
        percentage: 2.509350990731744,
      },
      {
        language: "YAML",
        percentage: 1.8476285081275758,
      },
      {
        language: "Markdown",
        percentage: 1.2951524748884733,
      },
      {
        language: "Docker",
        percentage: 1.0098753654467754,
      },
      {
        language: "Makefile",
        percentage: 0.9813351151858577,
      },
      {
        language: "Bash",
        percentage: 0.22905437541931054,
      },
      {
        language: "HTML",
        percentage: 0.17221410872998397,
      },
      {
        language: "INI",
        percentage: 0.09384441863674059,
      },
      {
        language: "TOML",
        percentage: 0.0857006528985198,
      },
      {
        language: "Prisma",
        percentage: 0.0415336352283311,
      },
      {
        language: "Git",
        percentage: 0.027438704444846543,
      },
      {
        language: "Git Config",
        percentage: 0.01736006980369162,
      },
      {
        language: "TSConfig",
        percentage: 0.013901128241566698,
      },
      {
        language: "CSS",
        percentage: 0.00989471300248098,
      },
    ],
  },
  projectStats: {
    durations: [
      {
        project: "cohere-toolkit",
        minutes: 690.3214306666666,
      },
      {
        project: "blobheart",
        minutes: 363.79199585000003,
      },
      {
        project: "guac",
        minutes: 58.79560266666667,
      },
      {
        project: "rbay",
        minutes: 53.37194244999999,
      },
      {
        project: "langchain-udemy-course",
        minutes: 24.214685049999996,
      },
      {
        project: "homebrew",
        minutes: 0.6909898666666667,
      },
    ],
  },
  startDate: "2024-07-14",
  endDate: "2024-07-20",
};
