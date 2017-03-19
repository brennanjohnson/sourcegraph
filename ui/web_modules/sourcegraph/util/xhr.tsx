import "whatwg-fetch";

import { context } from "sourcegraph/app/context";

// This file provides a common entrypoint to the fetch API.

export function combineHeaders(a: Headers, b: Headers): Headers {
	let headers = new Headers(a);
	b.forEach((val: string, name: any) => { headers.append(name, val); });
	return headers;
}

function getHeaders(init?: RequestInit): Headers | undefined {
	if (!(window as any).Headers) {
		// Avoid the Headers API if the browser is too hold to support it.
		return undefined;
	}
	let defaultHeaders = new Headers();
	Object.keys(context.xhrHeaders).forEach((key) => {
		defaultHeaders.set(key, context.xhrHeaders[key]);
	});
	return (init && init.headers) ? combineHeaders(defaultHeaders, new Headers(init.headers)) : defaultHeaders;
}
// defaultFetch wraps the fetch API.
//
// Note: the caller might wrap this with singleflightFetch.
export function defaultFetch(url: string | Request, init?: RequestInit): Promise<Response> {
	if (typeof url !== "string") { throw new Error("url must be a string (complex requests are not yet supported)"); }

	return fetch(url, {
		method: (init && init.method) || "GET",
		headers: getHeaders(init),
		body: init && init.body,
		mode: (init && init.mode) || "cors",
		redirect: (init && init.redirect) || "follow",
		credentials: (init && init.credentials) || "same-origin",
		cache: (init && init.cache) || "default",
	});
}

// checkStatus is intended to be chained in a fetch call. For example:
//   fetch(...).then(checkStatus) ...
export function checkStatus(resp: Response): Promise<Response> | Response {
	if (resp.status >= 200 && resp.status <= 299) { return resp; }
	return resp.text().then<Response>((body) => {
		if (typeof document === "undefined") {
			// Don't log in the browser because the devtools network inspector
			// makes it easy enough to see failed HTTP requests.
			console.error(`HTTP fetch failed with status ${resp.status} ${resp.statusText}: ${resp.url}: ${body}`);
		}
		let err: Error;
		try {
			err = new Error(resp.status.toString());
			(err as any).body = JSON.parse(body);
		} catch (error) {
			err = new Error(resp.statusText);
			(err as any).body = body;
			(err as any).response = { status: resp.status, statusText: resp.statusText, url: resp.url };
		}
		throw err;
	});
}
