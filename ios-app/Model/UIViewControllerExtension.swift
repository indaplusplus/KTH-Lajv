//
//  UIViewControllerExtension.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-17.
//

import UIKit

extension UIViewController {
    
    func showError(_ message: String) {
        let alert = UIAlertController(title: "Sorry", message: message, preferredStyle: .alert)
        let cancel = UIAlertAction(title: "Ok", style: .default, handler: nil)
        alert.addAction(cancel)
        present(alert, animated: true, completion: nil)
    }
    
}
