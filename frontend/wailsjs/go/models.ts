export namespace terminal {
	
	export class InputEvent {
	    in: string;
	    out: string;
	    err?: any;
	    dir: string;
	    // Go type: time
	    time: any;
	
	    static createFrom(source: any = {}) {
	        return new InputEvent(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.in = source["in"];
	        this.out = source["out"];
	        this.err = source["err"];
	        this.dir = source["dir"];
	        this.time = this.convertValues(source["time"], null);
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
	export class TerminalSession {
	    id: string;
	    input: InputEvent;
	    history: InputEvent[];
	    dir: string;
	    previousDir?: string;
	    ctx?: any;
	
	    static createFrom(source: any = {}) {
	        return new TerminalSession(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.input = this.convertValues(source["input"], InputEvent);
	        this.history = this.convertValues(source["history"], InputEvent);
	        this.dir = source["dir"];
	        this.previousDir = source["previousDir"];
	        this.ctx = source["ctx"];
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

