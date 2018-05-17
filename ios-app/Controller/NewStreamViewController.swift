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
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        titleTextField.text = ""
        courseTextField.text = ""
        lecturerTextField.text = ""
    }
    
    override func touchesBegan(_ touches: Set<UITouch>, with event: UIEvent?) {
        view.endEditing(true)
    }
    
    @IBAction func startStream(_ sender: UIButton) {
        view.endEditing(true)
        guard let title = titleTextField.text, !title.isEmpty
            , let course = courseTextField.text, !course.isEmpty
            , let lecturer = lecturerTextField.text, !lecturer.isEmpty else {
            let alert = UIAlertController(title: "Wait!", message: "Please fill in all the fields before starting a stream.", preferredStyle: .alert)
            let ok = UIAlertAction(title: "Ok", style: .cancel, handler: nil)
            alert.addAction(ok)
            present(alert, animated: true, completion: nil)
            return
        }
        
        DataManager.shared.createStream(title: title, course: course, lecturer: lecturer) { (success, error, data) in
            if !success || error != nil {
                self.showError("You can't create a stream right now, try again later.")
            }
            
            guard let streamKey = data as? String else {
                self.showError("You can't create a stream right now, try again later.")
                return
            }
            
            self.performSegue(withIdentifier: Constants.Storyboard.Segues.StartStreamSegue, sender: streamKey)
        }
    }
    
    fileprivate func additionalSetup() {
        startStreamButton.layer.cornerRadius = 4.0
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == Constants.Storyboard.Segues.StartStreamSegue {
            if let navVC = segue.destination as? UINavigationController, let destination = navVC.viewControllers.first as? MyStreamViewController,
                let streamKey = sender as? String {
                destination.streamKey = streamKey
            }
        }
    }
    
}
