import { Skeleton } from "@/components/ui/skeleton";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "./components/ui/card";

export function SkeletonChart() {
  return (
    <Card>
      <CardHeader>
        <CardTitle>
          <Skeleton className="h-4 md:w-[300px]" />
        </CardTitle>
        <CardDescription>
          <Skeleton className="h-4 md:w-[300px]" />
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="min-h-[200px] w-full">
          <div className="h-fit w-fit flex flex-col space-y-3 m-4">
            <Skeleton className="h-[250px] md:w-[400px] rounded-xl" />
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
