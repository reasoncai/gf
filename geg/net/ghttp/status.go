package main

import (
    "gitee.com/johng/gf/g"
    "gitee.com/johng/gf/g/net/ghttp"
)

func main() {
    s := g.Server()
    s.BindHandler("/", func(r *ghttp.Request){
        r.Response.Writeln("halo 世界！")
        r.Response.WriteStatus(404)
    })
    s.BindStatusHandler(404, func(r *ghttp.Request){
        r.Response.Writeln("This is customized 404 page")
    })
    s.SetPort(8199)
    s.Run()
}