// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {bindings} from '../models';
import {time} from '../models';

export function GetAllTimeUsage():Promise<Array<bindings.UsageInfo>>;

export function GetUsageBetweenDates(arg1:time.Time,arg2:time.Time):Promise<Array<bindings.UsageInfo>>;
