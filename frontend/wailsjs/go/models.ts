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

