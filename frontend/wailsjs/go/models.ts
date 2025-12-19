export namespace config {
	
	export class Config {
	    srcDirs: string[];
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.srcDirs = source["srcDirs"];
	    }
	}

}

