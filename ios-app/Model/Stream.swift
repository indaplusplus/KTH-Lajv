//
//  Stream.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import Foundation

class Stream {
    
    fileprivate var _id: Int
    fileprivate var _course: String
    fileprivate var _room: String
    fileprivate var _lecturer: String
    fileprivate var _streamer: String
    fileprivate var _name: String
    fileprivate var _date: String
    fileprivate var _vod: String
    fileprivate var _stream: String
    fileprivate var _hls: String
    
    init(id: Int, course: String, room: String, lecturer: String, name: String, streamer: String, date: String, vod: String, stream: String, hls: String) {
        _id = id
        _course = course
        _room = room
        _lecturer = lecturer
        _streamer = streamer
        _name = name
        _date = date
        _vod = vod
        _stream = stream
        _hls = hls
    }
    
    class func fakeData() -> [Stream] {
        return [
            Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "MDI", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Algorithms", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Envariabelanalys", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: ""),Stream(id: 239, course: "Programming", room: "D1", lecturer: "Ric Glassey", name: "Super cool lecture", streamer: "Adam J", date: "13:15", vod: "", stream: "", hls: "")
        ]
    }
    
    var id: Int {
        return _id
    }
    
    var course: String {
        return _course
    }
    
    var room: String {
        return _room
    }
    
    var lecturer: String {
        return _lecturer
    }
    
    var streamer: String {
        return _streamer
    }
    
    var name: String {
        return _name
    }
    
    var date: String {
        return _date
    }
    
    var vod: String {
        return _vod
    }
    
    var stream: String {
        return _stream
    }
    
    var hls: String {
        return _hls
    }
    
}
