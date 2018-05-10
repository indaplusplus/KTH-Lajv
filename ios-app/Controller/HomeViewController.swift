//
//  HomeViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit

class HomeViewController: UIViewController {

    @IBOutlet fileprivate weak var streamButton: UIButton!
    @IBOutlet fileprivate weak var streamButtonBackground: UIView!
    @IBOutlet fileprivate weak var tableView: UITableView!
    
    fileprivate var model = [Stream]()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        additionalSetup()
        getData()
    }
    
    @IBAction func newStream(_ sender: UIButton) {
        
    }
    
    fileprivate func additionalSetup() {
        tableView.delegate = self
        tableView.dataSource = self
        streamButtonBackground.layer.cornerRadius = 32.0
    }
    
    fileprivate func getData() {
        model = Stream.fakeData()
        tableView.reloadData()
    }
    
}

extension HomeViewController: UITableViewDelegate, UITableViewDataSource {
    
    func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }
    
    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return model.count
    }
    
    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCell(withIdentifier: Constants.Storyboard.TableCells.StreamTableViewCell, for: indexPath)
            as! StreamTableViewCell
        cell.config(stream: model[indexPath.row])
        return cell
    }
    
}



