export namespace bindings {
	
	export class ConfigDto {
	    applications: string[];
	    check_interval: string;
	
	    static createFrom(source: any = {}) {
	        return new ConfigDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.applications = source["applications"];
	        this.check_interval = source["check_interval"];
	    }
	}
	export class ProcessDto {
	    name: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	    }
	}
	export class ProcessListDto {
	    processes: ProcessDto[];
	
	    static createFrom(source: any = {}) {
	        return new ProcessListDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.processes = this.convertValues(source["processes"], ProcessDto);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProcessUsageData {
	    name: string;
	    seen: string;
	    duration: number;
	
	    static createFrom(source: any = {}) {
	        return new ProcessUsageData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.seen = source["seen"];
	        this.duration = source["duration"];
	    }
	}

}

