# parallel_crawler
Crawler that recursively scans and downloads the contents of a web site, preserving folder structure. Executes in parallel goroutines

# Launch parameters
To launch parallel crawler on https://yandex.ru with the maximum recursion depth of 4 levels, type:
./parallel_crawler.exe -max_depth=4 -url=https://yandex.ru
