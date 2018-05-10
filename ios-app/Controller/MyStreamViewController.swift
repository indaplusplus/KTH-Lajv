//
//  MyStreamViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-08.
//

import UIKit

class MyStreamViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
    }
    
    @IBAction func endStream(_ sender: UIBarButtonItem) {
        let alert = UIAlertController(title: "Are you sure?", message: "Do you really want to end the stream right now? All viewers will automatically be kicked out.", preferredStyle: .alert)
        let yes = UIAlertAction(title: "Yes", style: .cancel) { (_) in
            // End stream logic
            self.presentingViewController?.dismiss(animated: true, completion: nil)
        }
        let no = UIAlertAction(title: "No", style: .default, handler: nil)
        alert.addAction(yes)
        alert.addAction(no)
        present(alert, animated: true, completion: nil)
    }
    
    
}
