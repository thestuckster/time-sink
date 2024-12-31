export namespace bindings {
	
	export class ProcessUsageInfo {
	    name: string;
	    seen: string;
	    duration: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessUsageInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.seen = source["seen"];
	        this.duration = source["duration"];
	    }
	}

}

