//
//  LoginViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit

class LoginViewController: UIViewController {

    @IBOutlet fileprivate weak var loginButton: UIButton!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        additionalSetup()
    }
    
    fileprivate func additionalSetup() {
        loginButton.layer.cornerRadius = 4.0
    }
    
    
    @IBAction func login(_ sender: UIButton) {
        // TODO:- KTH API LOGIN AUTHENTICATION
        performSegue(withIdentifier: Constants.Storyboard.Segues.ShowFeedSegue, sender: nil)
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        if segue.identifier == Constants.Storyboard.Segues.ShowFeedSegue {
            if let destination = segue.destination as? HomeViewController {
                // TODO:- SETUP AFTER LOGIN STATE
            }
        }
    }
}
