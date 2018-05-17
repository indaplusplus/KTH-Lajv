//
//  DataManager.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-15.
//

import Foundation
import SwiftyJSON

typealias completed = (_ success: Bool, _ error: Error?, _ data: Any?) -> ()
class DataManager {
    // Singleton
    fileprivate static var _shared = DataManager()
    
    static var shared: DataManager {
        return _shared
    }
    
    func createStream(title: String, course: String, lecturer: String, completed: @escaping completed) {
        guard let url = URL(string: "http://live.edstrom.me:1339/stream/create") else {
            print("DataManager: Failed to find link")
            completed(false, nil, nil)
            return
        }
        
        let parameters = [
            "course" : course,
            "room" : "na",
            "streamer" : "Anonymous",
            "lecturer": lecturer,
            "name" : title,
            "token" : "",
        ]
        
        print("preparing to create stream with parameters: \(parameters)")
        
        var request = URLRequest(url: url)
        request.httpMethod = "POST"
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        request.addValue("application/json", forHTTPHeaderField: "Accept")
        request.httpBody = try! JSONSerialization.data(withJSONObject: parameters, options: .prettyPrinted)
        
        URLSession.shared.dataTask(with: request) { (data, response, error) in
            if error != nil {
                print("DataManager: Failed to create stream with error: \(error?.localizedDescription)")
                completed(false, error, nil)
                return
            }
            
            let json = JSON(data!)
            let streamKey = json["key"].stringValue
            completed(true, nil, streamKey)
            
        }.resume()
        
    }
    
    func getStreams(completion: @escaping completed) {
        guard let url = URL(string: "http://live.edstrom.me:1339/stream/all") else {
            print("DataManager: Failed to find link")
            completion(false, nil, nil)
            return
        }
        
        print("preparing to get data from url")
        var request = URLRequest(url: url)
        request.httpMethod = "GET"
        
        URLSession.shared.dataTask(with: request) { (data, response, error) in
            if error != nil {
                print("failed to get data error: \(error?.localizedDescription)")
                completion(false, error, nil)
                return
            }
            
            // Print data/result
            let json = JSON(data!)["data"]
            print(json)
            let streams = Stream.jsonToStream(json: json)
            completion(true, nil, streams)
            
        }.resume()
    }
}
