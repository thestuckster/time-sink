import {bindings} from "../../wailsjs/go/models";
import UsageInfo = bindings.UsageInfo;

export function convertDataSecondsToHours(data: UsageInfo[]) {


  data.forEach(item => {
    const rounded = Math.round(item.Duration / 3600)
    console.info("Rounded: " + rounded + " Raw: " + item.Duration / 3600);
    if (rounded > 0) {
      item.Duration = rounded;
    } else {
      item.Duration = item.Duration / 3600;
    }
  });

  return data;
}