<h1 align="center">
<sub>
<img  src="https://raw.githubusercontent.com/ViewTube/viewtube-vue/development/.github/images/logo.png"
      height="50"
      width="50">
</sub>
ViewTube
</h1>

<p align="center">
  <a href="https://github.com/ViewTube/viewtube-vue/actions/workflows/build-workflow.yml">
    <img src="https://github.com/ViewTube/viewtube-vue/actions/workflows/build-workflow.yml/badge.svg?branch=development" alt="Github Actions Status"> 
  </a>
  <a href="https://github.com/ViewTube/viewtube-vue/actions/workflows/test-workflow.yml">
    <img src="https://github.com/ViewTube/viewtube-vue/actions/workflows/test-workflow.yml/badge.svg?branch=development" alt="Daily Tests Status">
  </a>
  <a href="https://deepscan.io/dashboard#view=project&amp;tid=11097&amp;pid=14017&amp;bid=262917">
    <img src="https://deepscan.io/api/teams/11097/projects/14017/branches/262917/badge/grade.svg" alt="DeepScan grade">
  </a>
  <a href="https://github.com/ViewTube/viewtube-vue/blob/development/LICENSE">
    <img src="https://img.shields.io/github/license/mauriceoegerli/viewtube-vue" alt="GitHub">
  </a>
  <a href="https://hub.docker.com/r/mauriceo/viewtube">
    <img src="https://img.shields.io/docker/image-size/mauriceo/viewtube/dev?label=docker%20image" alt="Docker Image Size (tag)">
  </a>
  <img src="https://img.shields.io/github/last-commit/viewtube/viewtube-vue" alt="GitHub last commit">
</p>

ViewTube is an alternative YouTube frontend.

It can recommend, play and search for videos. It saves your watch progress and you can subscribe to channels for them to appear in your subscription feed.

It's built using [Nuxt.js](https://nuxtjs.org/) and [Nest](https://nestjs.com/).

Check [uptime.m-o.dev/status](https://uptime.m-o.dev/status) for updates and downtime status.  
Try the dev version on [dev.viewtube.io](https://dev.viewtube.io).

<h2 align="center">
<sub>
<img  src=".github/icons/star.svg"
      height="30"
      width="30">
</sub>
Features
</h2>

- Watch videos without ads or tracking
- Built from the ground up to be mobile and desktop friendly
- Dark and light themes
- Touch friendly video player with gestures
- Supports loop, speed, autoplay and volume
- Create an account separately from Youtube
- Read comments
- Search for videos
- Watch playlists
- Subscribe to channels and see their latest uploads
- Receive push notifications for subscribed channels
- Integrated SponsorBlock support

<h2 align="center">
<sub>
<img  src=".github/icons/question.svg"
      height="30"
      width="30">
</sub>
Where does ViewTube get the data from?
</h2>

ViewTube does not use the official Youtube API. It instead scrapes the data from the website using a combination of custom built tools and the following open source libraries.

- [node-ytdl-core](https://github.com/fent/node-ytdl-core)
- [node-ytsr](https://github.com/TimeForANinja/node-ytsr)
- [node-ytpl](https://github.com/TimeForANinja/node-ytpl)
- [yt-comment-scraper ](https://github.com/FreeTubeApp/yt-comment-scraper)

The [Invidious](https://github.com/iv-org/invidious) API is still used in a few places.
You can check the progress on [replacing the invidious api here](https://github.com/ViewTube/viewtube-vue/wiki/Invidious-API-migration).

<h2 align="center">
<sub>
<img  src=".github/icons/screenshot.svg"
      height="30"
      width="30">
</sub>
Screenshots
</h2>

<h3 align="center">
Homepage
</h3>

![Screenshot-Homepage](https://i.ibb.co/Gk5tKQ7/lxt1y0mk.jpg)

<h3 align="center">
Video
</h3>

![Screenshot-Video](https://i.ibb.co/RTL2v3f/g2ejf7wf.jpg)

<h3 align="center">
Channel
</h3>

![Screenshot-Channel](https://i.ibb.co/h9mf1yd/6j45ao5r.jpg)

<h2 align="center">
<sub>
<img  src=".github/icons/home.svg"
      height="30"
      width="30">
</sub>
Host it yourself
</h2>

A comprehensive installation guide is available in the [wiki](https://github.com/ViewTube/viewtube-vue/wiki/Installation)

<h2 align="center">
<sub>
<img  src=".github/icons/dev.svg"
      height="30"
      width="30">
</sub>
Development setup
</h2>
Prerequisites

- Nodejs v18.x
- Pnpm 7.3.x
- MongoDB 5.0.x

If you don't want to install locally, use the supplied docker-compose.dev.yml

- MongoDB 4.4.x
- Redis 6.x

Instructions

1. Clone this repo

   `git clone https://github.com/ViewTube/viewtube-vue`

2. Install dependencies

   `pnpm install`

   or use docker-compose to install more easily and hassle-free

   `pnpm run serve:env`

3. Start the components seperately in multiple terminals

   `pnpm run serve:server`  
   `pnpm run serve:client`

4. Visit [localhost:8066](http://localhost:8066) with a browser

Pull requests welcome 😄!

<h2 align="center">
<sub>
<img  src=".github/icons/people.svg"
      height="30"
      width="30">
</sub>
Contributors
</h2>

Thanks goes to these wonderful people

[emoji reference](https://allcontributors.org/docs/en/emoji-key)

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/cyacedev"><img src="https://avatars0.githubusercontent.com/u/46712905?v=4?s=100" width="100px;" alt=""/><br /><sub><b>cyacedev</b></sub></a><br /><a href="https://github.com/ViewTube/viewtube-vue/commits?author=cyacedev" title="Code">💻</a> <a href="https://github.com/ViewTube/viewtube-vue/commits?author=cyacedev" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/ckVendrix"><img src="https://avatars2.githubusercontent.com/u/51775140?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Vendrix</b></sub></a><br /><a href="https://github.com/ViewTube/viewtube-vue/commits?author=ckVendrix" title="Code">💻</a></td>
    <td align="center"><a href="https://github.com/Mennaruuk"><img src="https://avatars.githubusercontent.com/u/52135169?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Mennaruuk</b></sub></a><br /><a href="https://github.com/ViewTube/viewtube-vue/commits?author=Mennaruuk" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/mizzunet"><img src="https://avatars.githubusercontent.com/u/10193999?v=4?s=100" width="100px;" alt=""/><br /><sub><b>Missu</b></sub></a><br /><a href="https://github.com/ViewTube/viewtube-vue/commits?author=mizzunet" title="Documentation">📖</a></td>
    <td align="center"><a href="https://github.com/beardeddude"><img src="https://avatars.githubusercontent.com/u/36680638?v=4?s=100" width="100px;" alt=""/><br /><sub><b>beardeddude</b></sub></a><br /><a href="https://github.com/ViewTube/viewtube-vue/commits?author=beardeddude" title="Code">💻</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

# Mac Developer ENV
* Node.js v19.4.0 to /usr/local/bin/node
* npm v9.2.0 to /usr/local/bin/npm
* Make sure that /usr/local/bin is in your $PATH.

* /usr/local/bin/pnpm
