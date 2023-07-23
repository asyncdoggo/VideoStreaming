
function LoadVideo(video_id) {
    var video = document.getElementById('video');
    var videoSrc = `/api/video/${video_id}/video.m3u8`;
    if (Hls.isSupported()) {
        var hls = new Hls();
        hls.loadSource(videoSrc);
        hls.attachMedia(video);
    }

    else if (video.canPlayType('application/vnd.apple.mpegurl')) {
        video.src = videoSrc;
    }

    video.play()

}


document.getElementById("upload_button").addEventListener("click", async function () {
    const file = document.getElementById("video_upload").files[0]
    if (file != undefined) {
        var formdata = new FormData()
        formdata.append("file", file, file.name)
        const response = await fetch("/api/video/upload", {
            method: 'POST',
            body: formdata
        }).then((response) => response.json())

        if (response.status == "success") {
            document.getElementById("error").innerHTML = "uploaded"
            getVideos()
        }
    }
})


async function getVideos() {
    const list_div = document.getElementById("videoList")

    const response = await fetch("/api/videos", {
        method: "GET",
    })


    const data = await response.json()
    const files = data.files
    list_div.innerHTML = ""
    for (i in files) {
        list_div.innerHTML += `
        <p><a href="#" onClick=LoadVideo('${files[i].FileId}')>${files[i].name}<a/></p>
        `
    }

}

getVideos()