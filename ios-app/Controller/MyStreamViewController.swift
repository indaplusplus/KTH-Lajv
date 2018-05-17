//
//  MyStreamViewController.swift
//  kthlive
//
//  Created by Adam Jafer on 2018-05-08.
//

import UIKit
import AVKit
import AVFoundation
import HaishinKit
import VideoToolbox

class MyStreamViewController: UIViewController {
    
    fileprivate let session: AVAudioSession = AVAudioSession.sharedInstance()
    fileprivate let rtmpConnection:RTMPConnection = RTMPConnection()
    fileprivate var rtmpStream: RTMPStream!
    fileprivate var currentPosition: AVCaptureDevice.Position = .back
    
    var streamKey: String!
    
    override func viewDidLoad() {
        super.viewDidLoad()
        additionalSetup()
        setupAVSession()
        transparentNavBar()
    }
    
    fileprivate func additionalSetup() {
        addRightButton()
    }
    
    fileprivate func setupAVSession() {
        do {
            try self.session.setPreferredSampleRate(44_100)
            try self.session.setCategory(AVAudioSessionCategoryPlayAndRecord, with: .allowBluetooth)
            try self.session.setMode(AVAudioSessionModeDefault)
            try self.session.setActive(true)
            self.setupStream()
        } catch {
            let alert = UIAlertController(title: "Sorry",
                                          message: "Your camera isn't available right now, try again later",
                                          preferredStyle: .alert)
            let ok = UIAlertAction(title: "Ok", style: .cancel, handler: nil)
            alert.addAction(ok)
            self.present(alert, animated: true, completion: nil)
        }
    }
    
    fileprivate func transparentNavBar() {
        self.navigationController?.navigationBar.setBackgroundImage(UIImage(), for: .default)
        self.navigationController?.navigationBar.shadowImage = UIImage()
        self.navigationController?.navigationBar.isTranslucent = true
        self.navigationController?.view.backgroundColor = .clear
        self.navigationController?.navigationBar.tintColor = UIColor.white
    }
    
    fileprivate func setupStream() {
        rtmpStream = RTMPStream(connection: rtmpConnection)
        rtmpExtraSettings()
        rtmpStream.attachAudio(AVCaptureDevice.default(for: AVMediaType.audio)) { error in
            // print(error)
        }
        rtmpStream.attachCamera(DeviceUtil.device(withPosition: .back)) { error in
            print("f2 \(error.localizedDescription)")
            // print(error)
        }
        
        let lfView: LFView = LFView(frame: view.bounds)
        lfView.videoGravity = AVLayerVideoGravity.resizeAspectFill
        lfView.attachStream(rtmpStream)
        view.addSubview(lfView)
        
        rtmpConnection.connect("rtmp://live.edstrom.me/live")
        rtmpStream.publish(streamKey!)
    }
    
    fileprivate func rtmpExtraSettings() {
        rtmpStream.captureSettings = [
            "fps": 30,
            "sessionPreset": AVCaptureSession.Preset.medium,
            "continuousAutofocus": false,
            "continuousExposure": false
        ]
        rtmpStream.audioSettings = [
            "muted": false,
            "bitrate": 32 * 1024
        ]
        rtmpStream.videoSettings = [
            "width": 640,
            "height": 360,
            "bitrate": 160 * 1024,
            "profileLevel": kVTProfileLevel_H264_Baseline_3_1,
            "maxKeyFrameIntervalDuration": 2
        ]
        rtmpStream.recorderSettings = [
            AVMediaType.audio: [
                AVFormatIDKey: Int(kAudioFormatMPEG4AAC),
                AVSampleRateKey: 0,
                AVNumberOfChannelsKey: 0
            ],
            AVMediaType.video: [
                AVVideoCodecKey: AVVideoCodecH264,
                AVVideoHeightKey: 0,
                AVVideoWidthKey: 0
            ],
        ]
    }
    
    fileprivate func addRightButton() {
        let rightbut = UIBarButtonItem(image: #imageLiteral(resourceName: "camera"), style: .plain, target: self, action: #selector(turnCamera))
        navigationItem.setRightBarButton(rightbut, animated: true)
    }
    
    @objc fileprivate func turnCamera() {
        let position: AVCaptureDevice.Position = currentPosition == .back ? .front : .back
        rtmpStream.attachCamera(DeviceUtil.device(withPosition: position)) { error in
            // error turning camera
        }
        currentPosition = position
    }
    
    fileprivate func closeStream() {
        rtmpStream.close()
        rtmpStream.dispose()
    }
    
    @IBAction func endStream(_ sender: UIBarButtonItem) {
        let alert = UIAlertController(title: "Are you sure?", message: "Do you really want to end the stream right now? All viewers will automatically be kicked out.", preferredStyle: .alert)
        let yes = UIAlertAction(title: "Yes", style: .cancel) { (_) in
            self.closeStream()
            self.presentingViewController?.dismiss(animated: true, completion: nil)
        }
        let no = UIAlertAction(title: "No", style: .default, handler: nil)
        alert.addAction(yes)
        alert.addAction(no)
        present(alert, animated: true, completion: nil)
    }
}
