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
	export class ProcessUsageInfoDto {
	    name: string;
	    seen: string;
	    duration: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessUsageInfoDto(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.seen = source["seen"];
	        this.duration = source["duration"];
	    }
	}

}

