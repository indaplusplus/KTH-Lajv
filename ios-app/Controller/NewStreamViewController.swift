//
//  NewStreamViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-08.
//

import UIKit

class NewStreamViewController: UIViewController {
    
    @IBOutlet fileprivate weak var titleTextField: UITextField!
    @IBOutlet fileprivate weak var courseTextField: UITextField!
    @IBOutlet fileprivate weak var lecturerTextField: UITextField!
    @IBOutlet fileprivate weak var startStreamButton: UIButton!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        additionalSetup()
    }
    
    @IBAction func startStream(_ sender: UIButton) {
        guard let title = titleTextField.text, !title.isEmpty
            , let course = courseTextField.text, !course.isEmpty
            , let lecturer = lecturerTextField.text, !lecturer.isEmpty else {
            let alert = UIAlertController(title: "Wait!", message: "Please fill in all the fields before starting a stream.", preferredStyle: .alert)
            let ok = UIAlertAction(title: "Ok", style: .cancel, handler: nil)
            alert.addAction(ok)
            present(alert, animated: true, completion: nil)
            return
        }
        
        // Start stream logic
        performSegue(withIdentifier: Constants.Storyboard.Segues.StartStreamSegue, sender: nil)
    }
    
    fileprivate func additionalSetup() {
        startStreamButton.layer.cornerRadius = 4.0
    }
    
}
