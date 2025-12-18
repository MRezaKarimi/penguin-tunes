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

export namespace indexer {
	
	export class Track {
	    id: string;
	    path: string;
	    title: string;
	    album: string;
	    artist: string;
	    composer: string;
	    genre: string;
	    track_number: number;
	    cover: string;
	    year: number;
	
	    static createFrom(source: any = {}) {
	        return new Track(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.title = source["title"];
	        this.album = source["album"];
	        this.artist = source["artist"];
	        this.composer = source["composer"];
	        this.genre = source["genre"];
	        this.track_number = source["track_number"];
	        this.cover = source["cover"];
	        this.year = source["year"];
	    }
	}

}

export namespace main {
	
	export class Group {
	    name: string;
	    tracks: indexer.Track[];
	
	    static createFrom(source: any = {}) {
	        return new Group(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.tracks = this.convertValues(source["tracks"], indexer.Track);
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

}

