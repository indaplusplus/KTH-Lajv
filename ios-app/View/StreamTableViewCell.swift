//
//  StreamTableViewCell.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit

class StreamTableViewCell: UITableViewCell {
    @IBOutlet weak var thumbnailImage: UIImageView!
    @IBOutlet weak var thumbnailView: UIView!
    @IBOutlet weak var streamTitle: UILabel!
    @IBOutlet weak var streamCourse: UILabel!
    @IBOutlet weak var streamLecturer: UILabel!
    @IBOutlet weak var streamDate: UILabel!
    
    override func layoutIfNeeded() {
        super.layoutIfNeeded()
        thumbnailView.layer.cornerRadius = 4.0
        thumbnailImage.layer.cornerRadius = 4.0
    }
    
    func config(stream: Stream) {
        streamTitle.text = stream.name
        streamCourse.text = stream.course
        streamLecturer.text = stream.lecturer
        streamDate.text = stream.date
    }
}
