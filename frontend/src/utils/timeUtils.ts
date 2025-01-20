import {bindings} from "../../wailsjs/go/models";
import UsageInfo = bindings.UsageInfo;

export function convertDataSecondsToHours(data: UsageInfo[]) {
  data.forEach(item => {
    item.Duration = item.Duration / 3600;
  });

  return data;
}