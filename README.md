# tipmonkey

Let the monkey hold the cup so you can focus on making the music.

## High-level goals

Just collecting ideas, requirements, and other details here until they're more formally documented or implemented.

* Automatic conversion of tips/gifts to points (that can be used to request songs)
* Web and chatbot frontends for users to see queue, request or bump songs, etc.
* Broadcaster/mod-specific permissions for managing the queue, requests, and points
* API for connecting frontends to various backends (Spotify, Twitch, app DB)
* Reports for the broadcaster/mods that track statistics and trends
* Toggles for prevention of queue-hogging or request manipulation

## Technical/Development Details

### Points system

Everything boils down to points, which are used to determine queue priority. Points are assigned to users through tips/bits/subs/raids, and users use points to request or bump songs. The default point assignments are:

* $0.01 USD == 1 point (Twitch tip or direct tip)
* 1 bit == 1 point
* 1 sub == 250 points
* 20+ viewer raid == 1000 points
* 150+ viewer raid == 3000 points

Broadcasters can gift points to users, and if desired they can also allow mods the same privilege.

### Request system

Users can request songs using points, either through a chatbot or through a web UI. (More on that later.) A minimum point value can be set for standard requests, defaulting to 500 points. A separate minimum point value can be assigned to longer requests, defaulting to a length exceeding 5:30 and a point minimum of 750 points. The request system also has a couple additional features to prevent queue/request abuse and to encourage a wider set of users to participate.

The request system keeps track of how many times a user has requested a song in a given day/night/stream timeframe (defaulting to within the past 12 hours). Each successive request by a user within the timeframe is subjected to a point deduction, or a "greed tax". The default is -500 points per additional song, but can be set to a different value or a percentage. A "point floor" is also set to prevent requests from being added to the queue below a certain value, defaulting to 500.

The request system also checks to see if a user already has a song in the queue when a new request comes in. If so, it can auto-deduct points from that request. The default is to subtract 4000 points, but it can be set to a different value or percentage.

The request system can also do rounding or truncating of point values as an additional way to prevent queue/request manipulation. The default is not to, but if turned on the streamer can choose to round points up or down to the nearest threshold they set, like rounding up to the next 100 or down to the next 500.

#### Example one: The measured multi-requester

If the same user keeps dropping 2000 point requests, and they're waiting for their song to be played before putting in another, the point values of their requests in the queue would be: 2000, 1500, 1000, 500, 500, etc.

#### Example two: The greedy multi-requester

If the same user keeps dropping 2000 point requests, and they're dropping them all at once, the point values of their requests in the queue would be: 2000, 500, 500, etc.

#### Example three: The clever requester

Assuming it's turned on and set to round down to the nearest 250, if a user tried to jump ahead of the 500 point requests in the queue by submitting a request with 550 points, it would enter the queue at 500 points. Same with 700. If they requested at 750-999 points, it would enter at 750 points on the queue.

### Queue system

The queue table tracks current/unplayed requests. Each request is tracked with the following data:

* Request song title
* Request song artist
* Request song length
* Request song genre(s)
* Request song link*
* Request point value
* Requester/tipper username
* Request date/time* (stored as epoch, displayed as `YYYY/MM/DD HH:MM AM/PM`)
* Request "pre-tax" value* (points originally used to add it to the queue, before taxes, bumps, and other point manipulations)
* Request message (optional)*

Some pieces of data, marked with an asterisk above, are only viewable by the broadcaster and their mods.

### Chatbot

More info soon.

### Web UI

More info soon.

### Historical data

More info soon.
