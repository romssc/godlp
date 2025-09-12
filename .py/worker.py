#!/usr/bin/env python3
import sys, json, yt_dlp

class QuietLogger:
    def debug(self, msg): pass
    def warning(self, msg): pass
    def error(self, msg): pass

def dummy_progress_hook(d):
    pass

def fetch_metadata(url: str):
    ydl_opts = {
        "quiet": True,
        "no_warnings": True,
        "logger": QuietLogger(),
        "progress_hooks": [dummy_progress_hook],
    }
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=False)
    return {
        "title": info.get("title"),
        "duration": info.get("duration"),
        "uploader": info.get("uploader"),
        "upload_date": info.get("upload_date"),
        "thumbnails": info.get("thumbnails", []),
    }

def fetch_file(url: str, outtmpl: str, fmt: str):
    ydl_opts = {
        "outtmpl": outtmpl,
        "format": fmt,
        "quiet": True,
        "no_warnings": True,
        "logger": QuietLogger(),
        "progress_hooks": [dummy_progress_hook],
        "merge_output_format": "mp4",
        "postprocessors": [{
            "key": "FFmpegVideoConvertor",
            "preferedformat": "mp4",
        }],
    }
    with yt_dlp.YoutubeDL(ydl_opts) as ydl:
        info = ydl.extract_info(url, download=True)
    return {
        "path": ydl.prepare_filename(info), 
        "type": info.get("ext", "mp4")
    }

if __name__ == "__main__":
    if len(sys.argv) < 3:
        sys.stderr.write("USAGE: <worker.py> <command> <payload>\n")
        sys.exit(1)

    cmd = sys.argv[1]
    payload = json.loads(sys.argv[2])

    try:
        if cmd == "metadata":
            print(json.dumps(fetch_metadata(payload["url"])))
        elif cmd == "download":
            print(json.dumps(fetch_file(
                payload["url"],
                payload["output"],
                payload["format"],
            )))
        else:
            sys.stderr.write(f"UNKNOWN COMMAND: {cmd}\n")
            sys.exit(1)
    except Exception as e:
        sys.stderr.write(f"{e}\n")
        sys.exit(1)