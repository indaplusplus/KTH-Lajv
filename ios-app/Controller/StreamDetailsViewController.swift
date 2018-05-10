//
//  StreamDetailsViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit

class StreamDetailsViewController: UIViewController {
    
    @IBAction func options(_ sender: UIBarButtonItem) {
        let actionSheet = UIAlertController(title: "Options", message: "Choose from the list below", preferredStyle: .actionSheet)
        let action1 = UIAlertAction(title: "Option 1", style: .default) { (_) in
            // do something
        }
        let action2 = UIAlertAction(title: "Option 2", style: .default) { (_) in
            // do something
        }
        let action3 = UIAlertAction(title: "Option 3", style: .default) { (_) in
            // do something
        }
        let cancel = UIAlertAction(title: "Cancel", style: .cancel, handler: nil)
        actionSheet.addAction(action1)
        actionSheet.addAction(action2)
        actionSheet.addAction(action3)
        actionSheet.addAction(cancel)
        present(actionSheet, animated: true, completion: nil)
    }
}
