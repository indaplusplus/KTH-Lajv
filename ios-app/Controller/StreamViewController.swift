//
//  StreamViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit

class StreamViewController: UIViewController {

    override func viewDidLoad() {
        super.viewDidLoad()
        
    }
    
    @IBAction func showStreamDetails(_ sender: UIBarButtonItem) {
        performSegue(withIdentifier: Constants.Storyboard.Segues.ShowStreamInfoSegue, sender: nil)
    }
    
}
