import { context } from "sourcegraph/app/context";

class TelligentWrapper {
	private telligent: (...args: any[]) => void | null;
	private DEFAULT_ENV: string = "development";
	private PROD_ENV: string = "production";
	private DEFAULT_APP_ID: string = "UnknownApp";

	constructor() {
		if (global && global.window && global.window.telligent) {
			this.telligent = global.window.telligent;
		} else {
			return;
		}
		if (context.version !== "dev" && context.trackingAppID) {
			this.initialize(context.trackingAppID, this.PROD_ENV);
		} else {
			this.initialize(this.DEFAULT_APP_ID, this.DEFAULT_ENV);
		}
	}

	isTelligentLoaded(): boolean {
		return Boolean(this.telligent);
	}

	setUserId(loginInfo: string): void {
		if (!this.telligent) {
			return;
		}
		this.telligent("setUserId", loginInfo);
	}

	addStaticMetadataObject(metadata: any): void {
		if (!this.telligent) {
			return;
		}
		this.telligent("addStaticMetadataObject", metadata);
	}

	private addStaticMetadata(property: string, value: string, command: string): void {
		if (!this.telligent) {
			return;
		}
		this.telligent("addStaticMetadata", property, value, command);
	}

	setUserProperty(property: string, value: any): void {
		this.addStaticMetadata(property, value, "userInfo");
	}

	track(eventAction: string, eventProps: any): void {
		if (!this.telligent) {
			return;
		}
		this.telligent("track", eventAction, eventProps);
	}

	private initialize(appId: string, env: string): void {
		if (!this.telligent) {
			return;
		}
		this.telligent("newTracker", "sg", "sourcegraph-logging.telligentdata.com", {
			appId: appId,
			platform: "Web",
			encodeBase64: false,
			env: env,
			configUseCookies: true,
			useCookies: true,
			metadata: {
				gaCookies: true,
				performanceTiming: true,
				augurIdentityLite: true,
				webPage: true,
			},
		});
	}

}

export const telligent = new TelligentWrapper();
