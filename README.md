# GoFinder

Golang tool to parse and find specific data inside of a website's

## Installation

```
$ go get -u github.com/kkirsche/gofinder
```

## Usage

```
$ gofinder [options] [urls...]
```

## Example

```
~ ❯❯❯ gofinder -t https://www.verizon.com
INFO[0005] TITLE FOUND:	Verizon Fios & Custom TV | Internet, Cable & Phone
~ ❯❯❯ gofinder -c https://www.verizon.com
INFO[0000] COMMENT FOUND:	 VZ Tracker
INFO[0000] COMMENT FOUND:	<div class="legal content main wrap" id="disclaimer">
	*Speeds based on 75/75 Mbps. Download and upload time estimates based on maximum connection speeds. Actual speeds will vary. ^Speed and range tested in lab conditions using latest-generation wireless devices. In-home speed and range may vary depending on a number of factors like age/type of wireless devices used. Internet services, router location, distance and interference among other things. † Payment plans and prices vary based on number of connected TVs–up to 10. Media Servers and Media Clients beginning at $19.98/mo. Premium DVR Service is $32/mo. and Enhanced DVR service is $22/mo. Recorded content stored on current Set Top Box will not be transferred to new TV equipment. Specific model Verizon branded router required. FiOS Quantum TV available in select areas.</div>
INFO[0000] COMMENT FOUND:	 Modal
INFO[0000] COMMENT FOUND:	 Share
~ ❯❯❯ gofinder -l https://www.verizon.com
INFO[0000] LINK FOUND:	http://www.verizonwireless.com/?intcmp=VZ-GHOME-D-WIRELESS
INFO[0000] LINK FOUND:	/?lid=//global//residential
INFO[0000] LINK FOUND:	/home/verizonglobalhome/ghp_business.aspx
INFO[0000] LINK FOUND:	http://espanol.verizon.com/enes/home/verizonglobalhome/ghp_landing.aspx
INFO[0000] LINK FOUND:	http://www.verizonwireless.com/?intcmp=VZ-GHOME-D-WIRELESS
INFO[0000] LINK FOUND:	https://www.verizonwireless.com/smartphones/apple-iphone-x/
INFO[0000] LINK FOUND:	/?lid=//global//residential
INFO[0000] LINK FOUND:	/home/verizonglobalhome/ghp_business.aspx
INFO[0000] LINK FOUND:	/home/storelocator/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/careers/
INFO[0000] LINK FOUND:	javascript:O_LC();
INFO[0000] LINK FOUND:	http://www.verizon.com/about/privacy/
INFO[0000] LINK FOUND:	http://www.verizon.com/about/terms/
INFO[0000] LINK FOUND:	javascript:void(0);
INFO[0000] LINK FOUND:	javascript:void(0);
INFO[0000] LINK FOUND:	javascript:void(0);
```
