//
//  HomeViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-02.
//

import UIKit
import AVKit
import AVFoundation

class HomeViewController: UIViewController {

    @IBOutlet fileprivate weak var streamButton: UIButton!
    @IBOutlet fileprivate weak var streamButtonBackground: UIView!
    @IBOutlet fileprivate weak var tableView: UITableView!
    
    fileprivate let session: AVAudioSession = AVAudioSession.sharedInstance()
    fileprivate var model = [Stream]()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        additionalSetup()
        setupAVSession()
        getData()
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        self.navigationController?.navigationBar.tintColor = UIColor.darkGray
    }
    
    @IBAction func newStream(_ sender: UIButton) {
        
    }
    
    @IBAction func refresh(_ sender: UIBarButtonItem) {
        model = []
        tableView.reloadData()
        getData()
        print("refreshed")
    }
    
    fileprivate func additionalSetup() {
        tableView.delegate = self
        tableView.dataSource = self
        streamButtonBackground.layer.cornerRadius = 32.0
    }
    
    fileprivate func setupAVSession() {
        do {
            try session.setPreferredSampleRate(44_100)
            try session.setCategory(AVAudioSessionCategoryPlayAndRecord, with: .allowBluetooth)
            try session.setMode(AVAudioSessionModeDefault)
            try session.setActive(true)
        } catch {
            let alert = UIAlertController(title: "Sorry",
                                          message: "Video's aren't available right now, try again later.",
                                          preferredStyle: .alert)
            let ok = UIAlertAction(title: "Ok", style: .cancel, handler: nil)
            alert.addAction(ok)
            present(alert, animated: true, completion: nil)
        }
    }
    
    @objc fileprivate func getData() {
        DataManager.shared.getStreams { (success, error, data) in
            if !success || error != nil {
                self.showError("There are no livestreams available at the moment!")
                return
            }
            
            guard let streams = data as? [Stream] else {
                self.showError("There are no livestreams available at the moment!")
                return
            }
            
            DispatchQueue.main.async {
                self.model = streams
                self.tableView.reloadData()
            }
        }
    }
    
    fileprivate func transparentNavBar() {
        self.navigationController?.navigationBar.setBackgroundImage(UIImage(), for: .default)
        self.navigationController?.navigationBar.shadowImage = UIImage()
        self.navigationController?.navigationBar.isTranslucent = true
        self.navigationController?.view.backgroundColor = .clear
        self.navigationController?.navigationBar.tintColor = UIColor.white
    }
    
    fileprivate func addRightButton() {
        let barbut = UIBarButtonItem(image: #imageLiteral(resourceName: "more"), style: .plain, target: self, action: #selector(showMore))
        navigationItem.setRightBarButton(barbut, animated: true)
    }
    
    @objc fileprivate func showMore() {
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
    
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        let stream = model[indexPath.row]
        guard let url = URL(string: stream.hls) else {
            DispatchQueue.main.async {
                self.showError("This stream doesn't seem to be online right now, try again later")
            }
            return
        }
        
        let player = AVPlayer(url: url)
        let controller = AVPlayerViewController()
        // TODO:- Customize player so view controller auto plays video
        // add an extra view over the video to display comments
        // remove the pause button
        controller.player = player
        
        let barbut = UIBarButtonItem(image: #imageLiteral(resourceName: "more"), style: .plain, target: self, action: #selector(showMore))
        controller.navigationItem.setRightBarButton(barbut, animated: true)
        navigationController?.pushViewController(controller, animated: true)
        navigationController?.navigationItem.leftBarButtonItem?.customView?.alpha = 0.0
        transparentNavBar()
    }
    
}



