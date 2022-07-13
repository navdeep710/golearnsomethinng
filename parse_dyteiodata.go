package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"strings"
)

type TransportStats struct {
	BytesReceived   int64  `json:"bytesReceived"`
	BytesSent       int64  `json:"bytesSent"`
	DtlsCipher      string `json:"dtlsCipher"`
	DtlsState       string `json:"dtlsState"`
	IceRole         string `json:"iceRole"`
	PacketsReceived int64  `json:"packetsReceived"`
	PacketsSent     int64  `json:"packetsSent"`
	RoundTripTime   int64  `json:"roundTripTime"`
	StatId          string `json:"statId"`
	PeerId          string `json:"peerId"`
	EventId         string `json:"eventId"`
	TransportId     string `json:"transportId"`
	Consuming       bool   `json:"consuming"`
	Producing       bool   `json:"producing"`
}

func producingTransportStats(peerId string, eventId string, event gjson.Result) map[string][]byte {
	if event.Get("event").String() == "ping_stat" && event.Get("metaData.producingTransportStats").Exists() || event.Get("metaData.consumingTransportStats").Exists() {
		transportStats := TransportStats{
			BytesReceived: event.Get("metaData.producingTransportStats.stats.bytesReceived").Int(),
			BytesSent:     event.Get("metaData.producingTransportStats.stats.bytesSent").Int(),
			DtlsCipher:    event.Get("metaData.producingTransportStats.stats.dtlsCipher").String(),
			DtlsState:     event.Get("metaData.producingTransportStats.stats.dtlsState").String(),
			IceRole:       event.Get("metaData.producingTransportStats.stats.iceRole").String(),
			PacketsReceived: event.Get("metaData.producingTransportStats.sta	ts.packetsReceived").Int(),
			PacketsSent:   event.Get("metaData.producingTransportStats.stats.packetsSent").Int(),
			RoundTripTime: event.Get("metaData.producingTransportStats.stats.roundTripTime").Int(),
			StatId:        uuid.NewString(),
			PeerId:        peerId,
			EventId:       eventId,

			TransportId: event.Get("metaData.producingTransportStats.transportId").String(),
			Consuming:   event.Get("metaData.producingTransportStats.consuming").Bool(),
			Producing:   event.Get("metaData.producingTransportStats.producing").Bool(),
		}
		jsonBytesForKafka, _ := json.Marshal(&transportStats)
		transportStatMap := map[string][]byte{
			"transportStats": jsonBytesForKafka,
		}
		return transportStatMap
	}
	return make(map[string][]byte)
}

type Precallnetworkinformation struct {
	PeerId                string
	Location_latitude     float64
	Location_longitude    float64
	TurnConnectivity      bool
	HostConnectivity      bool
	RelayConnectivity     bool
	ReflexiveConnectivity bool
	EffectiveNetworkType  string
	Throughput            float64
	FractionalLoss        float64
	RTT                   float64
	Jitter                float64
	BackendRTT            float64
}

type ProducerAudioStats struct {
	StatId                               string  `json:"statId"`
	PeerId                               string  `json:"peerId"`
	EventId                              string  `json:"eventId"`
	ProducerId                           string  `json:"producerId"`
	BytesSent                            int64   `json:"bytesSent"`
	PacketsSent                          int64   `json:"packetsSent"`
	RetransmittedBytesSent               int64   `json:"retransmittedBytesSent"`
	RetransmittedPacketsSent             int64   `json:"retransmittedPacketsSent"`
	RemoteData_jitter                    float64 `json:"remoteData_jitter"`
	RemoteData_fractionLost              float64 `json:"remoteData_fractionLost"`
	RemoteData_roundTripTime             float64 `json:"remoteData_roundTripTime"`
	RemoteData_roundTripTimeMeasurements float64 `json:"remoteData_roundTripTimeMeasurements"`
	RemoteData_totalRoundTripTime        float64 `json:"remoteData_totalRoundTripTime"`
	RemoteData_packetsLost               int64   `json:"remoteData_packetsLost"`
}

type ProducerVideoStats struct {
	StatId  string `json:"statId"`
	PeerId  string `json:"peerId"`
	EventId string `json:"eventId"`

	ProducerId string `json:"producerId"`

	FrameWidth            float64 `json:"frameWidth"`
	FrameHeight           float64 `json:"frameHeight"`
	FramesEncoded         int64   `json:"framesEncoded"`
	FramesDropped         int64   `json:"framesDropped"`
	FramesPerSecond       int64   `json:"framesPerSecond"`
	FramesReceived        int64   `json:"framesReceived"`
	KeyFramesEncoded      int64   `json:"keyFramesEncoded"`
	FirCount              int64   `json:"firCount"`
	EncoderImplementation string  `json:"encoderImplementation"`

	BytesSent                int64 `json:"bytesSent"`
	PacketsSent              int64 `json:"packetsSent"`
	RetransmittedBytesSent   int64 `json:"retransmittedBytesSent"`
	RetransmittedPacketsSent int64 `json:"retransmittedPacketsSent"`

	HugeFramesSent                     int64  `json:"hugeFramesSent"`
	NackCount                          int64  `json:"nackCount"`
	PliCount                           int64  `json:"pliCount"`
	QpSum                              int64  `json:"qpSum"`
	QualityLimitationReason            string `json:"qualityLimitationReason"`
	QualityLimitationResolutionChanges int64  `json:"qualityLimitationResolutionChanges"`
	// type not known
	QualityLimitationDurations float64 `json:"qualityLimitationDurations"`
	TotalEncodeTime            int64   `json:"totalEncodeTime"`
	// type not known
	TotalPacketSendDelay                 int64   `json:"totalPacketSendDelay"`
	RemoteData_jitter                    float64 `json:"remoteData_jitter"`
	RemoteData_fractionLost              float64 `json:"remoteData_fractionLost"`
	RemoteData_roundTripTime             float64 `json:"remoteData_roundTripTime"`
	RemoteData_roundTripTimeMeasurements int64   `json:"remoteData_roundTripTimeMeasurements"`
	RemoteData_totalRoundTripTime        float64 `json:"remoteData_TotalRoundTripTime"`
	RemoteData_packetsLost               int64   `json:"remoteData_packetsLost"`
}

func getOrDefault(result gjson.Result, defaultValue gjson.Result) gjson.Result {
	if result.Exists() {
		return result
	}
	return defaultValue
}

func getOrDefaultIfParentExists(result gjson.Result, elementPath string, defaultValue gjson.Result) gjson.Result {
	// parent is the elementPath without the last element
	individualElements := strings.Split(elementPath, ".")
	parent := strings.Join(individualElements[:len(individualElements)-2], ".")
	if !result.Get(parent).Exists() {
		return defaultValue
	}
	return getOrDefault(result.Get(elementPath), defaultValue)
}

func processPreCallNetworkInformation(peerId string, eventId string, event gjson.Result) map[string][]byte {
	if event.Get("event").String() == "precall_end" {
		precallnetworkinformation := Precallnetworkinformation{
			PeerId:                peerId,
			Location_latitude:     getOrDefaultIfParentExists(event, "metaData.connectionInfo.location.coords.latitude", gjson.Result{}).Float(),
			Location_longitude:    getOrDefaultIfParentExists(event, "metaData.connectionInfo.location.coords.longitude", gjson.Result{}).Float(),
			TurnConnectivity:      event.Get("metaData.connectionInfo.turnConnectivity").Bool(),
			HostConnectivity:      getOrDefaultIfParentExists(event, "metaData.connectionInfo.connectivity.host", gjson.Result{}).Bool(),
			RelayConnectivity:     getOrDefaultIfParentExists(event, "metaData.connectionInfo.connectivity.relay", gjson.Result{}).Bool(),
			ReflexiveConnectivity: getOrDefaultIfParentExists(event, "metaData.connectionInfo.connectivity.reflexive", gjson.Result{}).Bool(),
			EffectiveNetworkType:  event.Get("metaData.connectionInfo.effectiveNetworkType").String(),
			Throughput:            event.Get("metaData.connectionInfo.throughput").Float(),
			FractionalLoss:        event.Get("metaData.connectionInfo.fractionalLoss").Float(),
			RTT:                   event.Get("metaData.connectionInfo.RTT").Float(),
			Jitter:                event.Get("metaData.connectionInfo.jitter").Float(),
			BackendRTT:            event.Get("metaData.connectionInfo.backendRTT").Float(),
		}
		jsonBytesForKafka, _ := json.Marshal(&precallnetworkinformation)
		precallnetworkinformationMap := map[string][]byte{
			"Precallnetworkinformation": jsonBytesForKafka,
		}
		return precallnetworkinformationMap
	}
	return make(map[string][]byte)
}

type IPInformation struct {
	PeerId   string `json:"peerId"`
	Ipv4     string `json:"ipv4"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func processIPInformation(peerId string, eventId string, event gjson.Result) map[string][]byte {
	if event.Get("event").String() == "precall_end" && event.Get("metaData.connectionInfo.ipDetails").Exists() {
		ipinformation := IPInformation{
			PeerId:   peerId,
			Ipv4:     event.Get("metaData.connectionInfo.ipDetails.ip").String(),
			City:     event.Get("metaData.connectionInfo.ipDetails.city").String(),
			Region:   event.Get("metaData.connectionInfo.ipDetails.region").String(),
			Country:  event.Get("metaData.connectionInfo.ipDetails.country").String(),
			Loc:      event.Get("metaData.connectionInfo.ipDetails.loc").String(),
			Org:      event.Get("metaData.connectionInfo.ipDetails.org").String(),
			Postal:   event.Get("metaData.connectionInfo.ipDetails.postal").String(),
			Timezone: event.Get("metaData.connectionInfo.ipDetails.timezone").String(),
		}
		jsonBytesForKafka, _ := json.Marshal(&ipinformation)
		ipInformationMap := map[string][]byte{
			"ipInformation": jsonBytesForKafka,
		}
		return ipInformationMap
	}

	return make(map[string][]byte)
}

type DeviceInfo struct {
	PeerId         string `json:"peerId"`
	UserAgent      string `json:"userAgent"`
	Cpus           int64  `json:"cpus"`
	Memory         int64  `json:"memory"`
	WebglSupport   bool   `json:"webglSupport"`
	IsMobile       bool   `json:"isMobile"`
	BrowserName    string `json:"browserName"`
	BrowserVersion string `json:"browserVersion"`
	OsName         string `json:"osName"`
	OsVersionName  string `json:"osVersionName"`
	EngineName     string `json:"engineName"`
}

func processDeviceInformation(peerId string, eventId string, event gjson.Result) map[string][]byte {
	if event.Get("event").String() == "call_join" && event.Get("metaData.peerMetaData.deviceInfo").Exists() {
		deviceinformation := DeviceInfo{
			PeerId:         peerId,
			UserAgent:      event.Get("metaData.peerMetaData.deviceInfo.userAgent").String(),
			Cpus:           event.Get("metaData.peerMetaData.deviceInfo.cpus").Int(),
			Memory:         event.Get("metaData.peerMetaData.deviceInfo.memory").Int(),
			WebglSupport:   event.Get("metaData.peerMetaData.deviceInfo.webglSupport").Bool(),
			IsMobile:       event.Get("metaData.peerMetaData.deviceInfo.isMobile").Bool(),
			BrowserName:    event.Get("metaData.peerMetaData.deviceInfo.browserName").String(),
			BrowserVersion: event.Get("metaData.peerMetaData.deviceInfo.browserVersion").String(),
			OsName:         event.Get("metaData.peerMetaData.deviceInfo.osName").String(),
			OsVersionName:  event.Get("metaData.peerMetaData.deviceInfo.osVersionName").String(),
			EngineName:     event.Get("metaData.peerMetaData.deviceInfo.engineName").String(),
		}
		jsonBytesForKafka, _ := json.Marshal(&deviceinformation)
		deviceInformationMap := map[string][]byte{
			"deviceInformation": jsonBytesForKafka,
		}
		return deviceInformationMap
	}
	return make(map[string][]byte)
}

func getProducerAudioStatsFromAudioStatsEntry(peerId string, eventId string, audioStats gjson.Result) []ProducerAudioStats {
	producerAudioStatsList := []ProducerAudioStats{}
	for _, audioStatsEntry := range audioStats.Get("audioStats").Array() {
		producerAudioStats := ProducerAudioStats{
			StatId:  uuid.NewString(),
			PeerId:  peerId,
			EventId: eventId,

			ProducerId: audioStats.Get("producerId").String(),

			BytesSent:                audioStatsEntry.Get("bytesSent").Int(),
			PacketsSent:              audioStatsEntry.Get("packetsSent").Int(),
			RetransmittedBytesSent:   audioStatsEntry.Get("retransmittedBytesSent").Int(),
			RetransmittedPacketsSent: audioStatsEntry.Get("retransmittedPacketsSent").Int(),
		}
		if audioStatsEntry.Get("remoteData").Exists() {
			producerAudioStats.RemoteData_jitter = audioStatsEntry.Get("remoteData.jitter").Float()
			producerAudioStats.RemoteData_fractionLost = audioStatsEntry.Get("remoteData.fractionLost").Float()
			producerAudioStats.RemoteData_roundTripTime = audioStatsEntry.Get("remoteData.roundTripTime").Float()
			producerAudioStats.RemoteData_roundTripTimeMeasurements = audioStatsEntry.Get("remoteData.roundTripTimeMeasurements").Float()
			producerAudioStats.RemoteData_totalRoundTripTime = audioStatsEntry.Get("remoteData.totalRoundTripTime").Float()
			producerAudioStats.RemoteData_packetsLost = audioStatsEntry.Get("remoteData.packetsLost").Int()
		}
		producerAudioStatsList = append(producerAudioStatsList, producerAudioStats)
	}
	return producerAudioStatsList
}

func processProducerAudioStatsEntries(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	if event.Get("event").String() == "ping_stat" && event.Get("metaData.producerStats").Exists() {
		producerAudioStatsArray := []map[string][]byte{}
		for _, producerStats := range event.Get("metaData.producerStats").Array() {
			producerAudioStats := getProducerAudioStatsFromAudioStatsEntry(peerId, eventId, producerStats)
			for _, producerAudioStats := range producerAudioStats {
				jsonBytesForKafka, _ := json.Marshal(&producerAudioStats)
				producerAudioStatsMap := map[string][]byte{
					"producerAudioStats": jsonBytesForKafka,
				}
				producerAudioStatsArray = append(producerAudioStatsArray, producerAudioStatsMap)
			}
		}
		return producerAudioStatsArray
	}
	return make([]map[string][]byte, 1)
}

func getProducerVideoStatsFromVideoStatsEntry(peerId string, eventId string, videoStats gjson.Result) []ProducerVideoStats {
	producerVideoStatsList := []ProducerVideoStats{}
	for _, videoStatsEntry := range videoStats.Get("videoStats").Array() {
		producerVideoStats := ProducerVideoStats{
			StatId:  uuid.NewString(),
			PeerId:  peerId,
			EventId: eventId,

			ProducerId: videoStats.Get("producerId").String(),

			FrameWidth:            videoStatsEntry.Get("frameWidth").Float(),
			FrameHeight:           videoStatsEntry.Get("frameHeight").Float(),
			FramesEncoded:         videoStatsEntry.Get("framesEncoded").Int(),
			FramesDropped:         videoStatsEntry.Get("framesDropped").Int(),
			FramesPerSecond:       videoStatsEntry.Get("framesPerSecond").Int(),
			FramesReceived:        videoStatsEntry.Get("framesReceived").Int(),
			KeyFramesEncoded:      videoStatsEntry.Get("keyFramesEncoded").Int(),
			FirCount:              videoStatsEntry.Get("firCount").Int(),
			EncoderImplementation: videoStatsEntry.Get("encoderImplementation").String(),

			BytesSent:                videoStatsEntry.Get("bytesSent").Int(),
			PacketsSent:              videoStatsEntry.Get("packetsSent").Int(),
			RetransmittedBytesSent:   videoStatsEntry.Get("retransmittedBytesSent").Int(),
			RetransmittedPacketsSent: videoStatsEntry.Get("retransmittedPacketsSent").Int(),

			HugeFramesSent:                     videoStatsEntry.Get("hugeFramesSent").Int(),
			NackCount:                          videoStatsEntry.Get("nackCount").Int(),
			PliCount:                           videoStatsEntry.Get("pliCount").Int(),
			QpSum:                              videoStatsEntry.Get("qpSum").Int(),
			QualityLimitationReason:            videoStatsEntry.Get("qualityLimitationReason").String(),
			QualityLimitationResolutionChanges: videoStatsEntry.Get("qualityLimitationResolutionChanges").Int(),
			QualityLimitationDurations:         videoStatsEntry.Get("qualityLimitationDurations").Float(),
			TotalEncodeTime:                    videoStatsEntry.Get("totalEncodeTime").Int(),
			TotalPacketSendDelay:               videoStatsEntry.Get("totalPacketSendDelay").Int(),
		}
		if videoStatsEntry.Get("remoteData").Exists() {
			producerVideoStats.RemoteData_jitter = videoStatsEntry.Get("remoteData.jitter").Float()
			producerVideoStats.RemoteData_fractionLost = videoStatsEntry.Get("remoteData.fractionLost").Float()
			producerVideoStats.RemoteData_roundTripTime = videoStatsEntry.Get("remoteData.roundTripTime").Float()
			producerVideoStats.RemoteData_roundTripTimeMeasurements = videoStatsEntry.Get("remoteData.roundTripTimeMeasurements").Int()
			producerVideoStats.RemoteData_totalRoundTripTime = videoStatsEntry.Get("remoteData.totalRoundTripTime").Float()
			producerVideoStats.RemoteData_packetsLost = videoStatsEntry.Get("remoteData.packetsLost").Int()
		}
		producerVideoStatsList = append(producerVideoStatsList, producerVideoStats)
	}
	return producerVideoStatsList
}

func processProducerVideoStatsEntries(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	if event.Get("event").String() == "ping_stat" && event.Get("metaData.producerStats").Exists() {
		producerVideoStatsArray := []map[string][]byte{}
		for _, producerStats := range event.Get("metaData.producerStats").Array() {
			producerVideoStats := getProducerVideoStatsFromVideoStatsEntry(peerId, eventId, producerStats)
			for _, producerVideoStats := range producerVideoStats {
				jsonBytesForKafka, _ := json.Marshal(&producerVideoStats)
				producerVideoStatsMap := map[string][]byte{
					"producerVideoStats": jsonBytesForKafka,
				}
				producerVideoStatsArray = append(producerVideoStatsArray, producerVideoStatsMap)
			}
		}
		return producerVideoStatsArray
	}
	return make([]map[string][]byte, 1)
}

type ConsumerAudioStats struct {
	// typo ??
	StatId  string `json:"statId"`
	PeerId  string `json:"peerId"`
	EventId string `json:"eventId"`

	ProducerId      string `json:"producerId"`
	ConsumerId      string `json:"consumerId"`
	ConsumingPeerId string `json:"consumingPeerId"`

	BytesReceived   int64   `json:"bytesReceived"`
	PacketsReceived int64   `json:"packetsReceived"`
	PacketsLost     int64   `json:"packetsLost"`
	Jitter          float64 `json:"jitter"`

	NackCount int64 `json:"nackCount"`

	AudioLevel               float64 `json:"audioLevel"`
	ConcealedSamples         int64   `json:"concealedSamples"`
	ConcealmentEvents        int64   `json:"concealmentEvents"`
	JitterBufferDelay        float64 `json:"jitterBufferDelay"`
	JitterBufferEmittedCount int64   `json:"jitterBufferEmittedCount"`
	TotalAudioEnergy         float64 `json:"totalAudioEnergy"`
	TotalSamplesDuration     float64 `json:"totalSamplesDuration"`
	TotalSamplesReceived     float64 `json:"totalSamplesReceived"`
}

func processConsumerAudioStatsEntry(peerId string, eventId string, videoStats gjson.Result) []ConsumerAudioStats {
	consumerAudioStatList := []ConsumerAudioStats{}
	for _, consumerAudioStatsEntry := range videoStats.Get("audioStats").Array() {
		consumerAudioStats := ConsumerAudioStats{
			StatId:          uuid.NewString(),
			PeerId:          peerId,
			EventId:         eventId,
			ProducerId:      videoStats.Get("producerId").String(),
			ConsumerId:      videoStats.Get("consumerId").String(),
			ConsumingPeerId: videoStats.Get("consumingPeerId").String(),

			BytesReceived:   consumerAudioStatsEntry.Get("bytesReceived").Int(),
			PacketsReceived: consumerAudioStatsEntry.Get("packetsReceived").Int(),
			PacketsLost:     consumerAudioStatsEntry.Get("packetsLost").Int(),
			Jitter:          consumerAudioStatsEntry.Get("jitter").Float(),

			NackCount: consumerAudioStatsEntry.Get("nackCount").Int(),

			AudioLevel:               consumerAudioStatsEntry.Get("audioLevel").Float(),
			ConcealedSamples:         consumerAudioStatsEntry.Get("concealedSamples").Int(),
			ConcealmentEvents:        consumerAudioStatsEntry.Get("concealmentEvents").Int(),
			JitterBufferDelay:        consumerAudioStatsEntry.Get("jitterBufferDelay").Float(),
			JitterBufferEmittedCount: consumerAudioStatsEntry.Get("jitterBufferEmittedCount").Int(),
			TotalAudioEnergy:         consumerAudioStatsEntry.Get("totalAudioEnergy").Float(),
			TotalSamplesDuration:     consumerAudioStatsEntry.Get("totalSamplesDuration").Float(),
			TotalSamplesReceived:     consumerAudioStatsEntry.Get("totalSamplesReceived").Float(),
		}
		consumerAudioStatList = append(consumerAudioStatList, consumerAudioStats)
	}
	return consumerAudioStatList
}

func processConsumerAudioStatsEntries(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	consumerAudioStatsArray := []map[string][]byte{}
	if event.Get("event").String() == "ping_stat" && event.Get("metaData.consumerStats").Exists() {

		for _, consumerStats := range event.Get("metaData.consumerStats").Array() {
			for _, consumerAudioStats := range processConsumerAudioStatsEntry(peerId, eventId, consumerStats) {
				jsonBytesForKafka, _ := json.Marshal(&consumerAudioStats)
				producerVideoStatsMap := map[string][]byte{
					"consumerAudioStats": jsonBytesForKafka,
				}
				consumerAudioStatsArray = append(consumerAudioStatsArray, producerVideoStatsMap)
			}
		}
	}
	return consumerAudioStatsArray
}

type ConsumerVideoStats struct {
	StatId  string `json:"statId"`
	PeerId  string `json:"peerId"`
	EventId string `json:"eventId"`

	ProducerId      string `json:"producerId"`
	ConsumerId      string `json:"consumerId"`
	ConsumingPeerId string `json:"consumingPeerId"`

	FrameWidth            int64  `json:"frameWidth"`
	FrameHeight           int64  `json:"frameHeight"`
	FramesDecoded         int64  `json:"framesDecoded"`
	FramesDropped         int64  `json:"framesDropped"`
	FramesPerSecond       int64  `json:"framesPerSecond"`
	FramesReceived        int64  `json:"framesReceived"`
	KeyFramesDecoded      int64  `json:"keyFramesDecoded"`
	FirCount              int64  `json:"firCount"`
	DecoderImplementation string `json:"decoderImplementation"`

	BytesReceived   int64   `json:"bytesReceived"`
	PacketsReceived int64   `json:"packetsReceived"`
	PacketsLost     int64   `json:"packetsLost"`
	Jitter          float64 `json:"jitter"`
}

func processConsumerVideoStatsEntry(peerId string, eventId string, videoStats gjson.Result) []ConsumerVideoStats {
	consumerVideoStatList := []ConsumerVideoStats{}
	for _, consumerVideoStatsEntry := range videoStats.Get("videoStats").Array() {
		consumerVideoStats := ConsumerVideoStats{
			StatId:          uuid.NewString(),
			PeerId:          peerId,
			EventId:         eventId,
			ProducerId:      videoStats.Get("producerId").String(),
			ConsumerId:      videoStats.Get("consumerId").String(),
			ConsumingPeerId: videoStats.Get("consumingPeerId").String(),

			FrameWidth:            consumerVideoStatsEntry.Get("frameWidth").Int(),
			FrameHeight:           consumerVideoStatsEntry.Get("frameHeight").Int(),
			FramesDecoded:         consumerVideoStatsEntry.Get("framesDecoded").Int(),
			FramesDropped:         consumerVideoStatsEntry.Get("framesDropped").Int(),
			FramesPerSecond:       consumerVideoStatsEntry.Get("framesPerSecond").Int(),
			FramesReceived:        consumerVideoStatsEntry.Get("framesReceived").Int(),
			KeyFramesDecoded:      consumerVideoStatsEntry.Get("keyFramesDecoded").Int(),
			FirCount:              consumerVideoStatsEntry.Get("firCount").Int(),
			DecoderImplementation: consumerVideoStatsEntry.Get("decoderImplementation").String(),

			BytesReceived:   consumerVideoStatsEntry.Get("bytesReceived").Int(),
			PacketsReceived: consumerVideoStatsEntry.Get("packetsReceived").Int(),
			PacketsLost:     consumerVideoStatsEntry.Get("packetsLost").Int(),
			Jitter:          consumerVideoStatsEntry.Get("jitter").Float(),
		}
		consumerVideoStatList = append(consumerVideoStatList, consumerVideoStats)
	}
	return consumerVideoStatList
}

func processConsumerVideoStatsEntries(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	consumerVideoStatsArray := []map[string][]byte{}
	if event.Get("event").String() == "ping_stat" && event.Get("metaData.consumerStats").Exists() {

		for _, consumerStats := range event.Get("metaData.consumerStats").Array() {
			for _, consumerVideoStats := range processConsumerVideoStatsEntry(peerId, eventId, consumerStats) {
				jsonBytesForKafka, _ := json.Marshal(&consumerVideoStats)
				producerVideoStatsMap := map[string][]byte{
					"consumerVideoStats": jsonBytesForKafka,
				}
				consumerVideoStatsArray = append(consumerVideoStatsArray, producerVideoStatsMap)
			}
		}
	}
	return consumerVideoStatsArray
}

type PeerInformation struct {
	PeerId      string `json:"peerId"`
	MeetingEnv  string `json:"meetingEnv"`
	Timestamp   string `json:"timestamp"`
	RoomName    string `json:"roomName"`
	RoomUUID    string `json:"roomUUID"`
	DisplayName string `json:"displayName"`
	UserId      string `json:"userId"`
	// we are storing it as string but will be available to be queried as map
	Permissions      string `json:"permissions"`
	RoomViewType     string `json:"roomViewType"`
	ClientSpecificId string `json:"clientSpecificId"`
	ParticipantRole  string `json:"participantRole"`
}

func processPeerInformation(peerId string, eventId string, event gjson.Result) map[string][]byte {
	if event.Get("event").String() == "call_join" {
		peerinformation := PeerInformation{
			PeerId:           peerId,
			MeetingEnv:       event.Get("metaData.peerMetaData.meetingEnv").String(),
			Timestamp:        event.Get("metaData.peerMetaData.timestamp").String(),
			RoomName:         event.Get("metaData.peerMetaData.roomName").String(),
			RoomUUID:         event.Get("metaData.peerMetaData.roomUUID").String(),
			DisplayName:      event.Get("metaData.peerMetaData.displayName").String(),
			UserId:           event.Get("metaData.peerMetaData.userId").String(),
			Permissions:      event.Get("metaData.peerMetaData.permissions").String(),
			RoomViewType:     event.Get("metaData.peerMetaData.roomViewType").String(),
			ClientSpecificId: event.Get("metaData.peerMetaData.clientSpecificId").String(),
			ParticipantRole:  event.Get("metaData.peerMetaData.participantRole").String(),
		}
		jsonBytesForKafka, _ := json.Marshal(&peerinformation)
		peerInformationMap := map[string][]byte{
			"peerInformation": jsonBytesForKafka,
		}
		return peerInformationMap
	}
	return make(map[string][]byte)
}

type DevicePermissions struct {
	DevicePermissionId string `json:"devicePermissionId"`
	DeviceType         string `json:"deviceType"`
	PermissionsGranted bool   `json:"permissionsGranted"`
	PeerId             string `json:"peerId"`
}

func processDevicePermissions(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	devicePermissionsArray := []map[string][]byte{}
	if event.Get("event").String() == "camera_permission" && event.Get("metaData.deviceType").Exists() && event.Get("metaData.permissions").Exists() && len(event.Get("metaData.permissions").Array()) == 1 {
		devicePermissions := DevicePermissions{
			DevicePermissionId: uuid.NewString(),
			DeviceType:         event.Get("metaData.deviceType").String(),
			PermissionsGranted: true,
			PeerId:             peerId,
		}
		jsonBytesForKafka, _ := json.Marshal(&devicePermissions)
		devicePermissionsMap := map[string][]byte{
			"devicePermissions": jsonBytesForKafka,
		}
		devicePermissionsArray = append(devicePermissionsArray, devicePermissionsMap)
	}

	return devicePermissionsArray
}

type Devices struct {
	DeviceId   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
	DeviceName string `json:"deviceName"`
	PeerId     string `json:"peerId"`
}

func processAudioDevices(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	devicesArray := []map[string][]byte{}
	if event.Get("event").String() == "audio_devices_updates" && event.Get("metaData.deviceList").Exists() {
		for _, device := range event.Get("metaData.deviceList").Array() {
			devices := Devices{
				DeviceId:   uuid.New().String(),
				DeviceType: "AUDIO",
				DeviceName: device.Get("label").String(),
				PeerId:     peerId,
			}
			jsonBytesForKafka, _ := json.Marshal(&devices)
			devicesMap := map[string][]byte{
				"devices": jsonBytesForKafka,
			}
			devicesArray = append(devicesArray, devicesMap)
		}
	}
	return devicesArray
}

func processVideoDevices(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	devicesArray := []map[string][]byte{}
	if event.Get("event").String() == "video_devices_updates" && event.Get("metaData.deviceList").Exists() {
		for _, device := range event.Get("metaData.deviceList").Array() {
			devices := Devices{
				DeviceId:   uuid.New().String(),
				DeviceType: "VIDEO",
				DeviceName: device.Get("label").String(),
				PeerId:     peerId,
			}
			jsonBytesForKafka, _ := json.Marshal(&devices)
			devicesMap := map[string][]byte{
				"devices": jsonBytesForKafka,
			}
			devicesArray = append(devicesArray, devicesMap)
		}
	}
	return devicesArray
}

func processSpeakerDevices(peerId string, eventId string, event gjson.Result) []map[string][]byte {
	devicesArray := []map[string][]byte{}
	if event.Get("event").String() == "speaker_devices_updates" && event.Get("metaData.deviceList").Exists() {
		for _, device := range event.Get("metaData.deviceList").Array() {
			devices := Devices{
				DeviceId:   uuid.New().String(),
				DeviceType: "SPEAKER",
				DeviceName: device.Get("label").String(),
				PeerId:     peerId,
			}
			jsonBytesForKafka, _ := json.Marshal(&devices)
			devicesMap := map[string][]byte{
				"devices": jsonBytesForKafka,
			}
			devicesArray = append(devicesArray, devicesMap)
		}
	}
	return devicesArray
}

func filterSliceForNonNull(slice []map[string][]byte) []map[string][]byte {
	var nonNullSlice []map[string][]byte
	for _, item := range slice {
		if len(item) > 0 {
			nonNullSlice = append(nonNullSlice, item)
		}
	}
	return nonNullSlice
}

func processEntries(peerId string, payload gjson.Result) []map[string][]byte {
	var entriesArray []map[string][]byte
	for _, v := range payload.Get("entries").Array() {
		eventId := uuid.New().String()
		transportStats := producingTransportStats(peerId, eventId, v)
		precallnetworkinformation := processPreCallNetworkInformation(peerId, eventId, v)
		ipinformation := processIPInformation(peerId, eventId, v)
		peerInformation := processPeerInformation(peerId, eventId, v)
		deviceInformation := processDeviceInformation(peerId, eventId, v)
		// this is list of producerAudioStats
		audioStatsInformation := processProducerAudioStatsEntries(peerId, eventId, v)
		// this is list of producerVideoStats
		producerVideoStats := processProducerVideoStatsEntries(peerId, eventId, v)
		consumerAudioStats := processConsumerAudioStatsEntries(peerId, eventId, v)
		consumerVideoStats := processConsumerVideoStatsEntries(peerId, eventId, v)
		devicePermissions := processDevicePermissions(peerId, eventId, v)
		audioDevices := processAudioDevices(peerId, eventId, v)
		videoDevices := processVideoDevices(peerId, eventId, v)
		speakerDevices := processSpeakerDevices(peerId, eventId, v)
		entriesArray = append(entriesArray)
		entriesArray = append(entriesArray, transportStats, precallnetworkinformation, ipinformation, peerInformation, deviceInformation)
		entriesArray = append(entriesArray, audioStatsInformation...)
		entriesArray = append(entriesArray, producerVideoStats...)
		entriesArray = append(entriesArray, consumerAudioStats...)
		entriesArray = append(entriesArray, consumerVideoStats...)
		entriesArray = append(entriesArray, devicePermissions...)
		entriesArray = append(entriesArray, audioDevices...)
		entriesArray = append(entriesArray, videoDevices...)
		entriesArray = append(entriesArray, speakerDevices...)
	}

	return filterSliceForNonNull(entriesArray)
}

func GetValuesFromDytePayload() {
	//str := `{"payload":{"entries":[{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":53060,"bytesSent":8056292,"packetsSent":11147,"packetsReceived":846,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.038,"totalRoundTripTime":0.492},"transportId":"4a7e5e27-1ba1-41b4-a386-7b6ed5f116fe","consuming":false,"producing":true},"consumingTransportStats":{"transportId":"a375600d-e910-4c17-a03e-33b93fd810a2","consuming":true,"producing":false},"producerStats":[{"producerId":"a6d0931d-efff-4b77-aacc-6264af4aa9f1","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":113509,"packetsSent":3149,"nackCount":0,"remoteData":{"jitter":0,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":12,"totalRoundTripTime":0.492,"packetsLost":2}}]},{"producerId":"c092ffa0-5ad1-4074-95db-84b35f8363b3","videoStats":[{"frameHeight":240,"frameWidth":320,"framesEncoded":1509,"framesPerSecond":24,"framesSent":1509,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":1,"pliCount":0,"qpSum":71294,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":63.045},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":6.795,"totalPacketSendDelay":6.795,"retransmittedBytesSent":4212,"retransmittedPacketsSent":6,"bytesSent":1052863,"packetsSent":1551,"nackCount":30,"remoteData":{"jitter":0.000044444444444444447,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":49,"totalRoundTripTime":2.101,"packetsLost":8}},{"frameHeight":480,"frameWidth":640,"framesEncoded":1509,"framesPerSecond":24,"framesSent":1509,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":0,"pliCount":0,"qpSum":34681,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":63.045},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":6.75,"totalPacketSendDelay":6.75,"retransmittedBytesSent":58894,"retransmittedPacketsSent":56,"bytesSent":6420855,"packetsSent":6232,"nackCount":94,"remoteData":{"jitter":0.000033333333333333335,"fractionLost":0,"roundTripTime":0.037,"roundTripTimeMeasurements":49,"totalRoundTripTime":2.153,"packetsLost":51}}],"audioStats":[]}],"consumerStats":[]},"timestamp":"2022-07-11T05:00:22.701Z"}]},"peerId":"8e94afa8-fcbe-4ed3-b3ce-1e040fec477e"}`
	//videostatsstr := `{"payload":{"entries":[{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":16797,"bytesSent":2690787,"packetsSent":3754,"packetsReceived":250,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.039,"totalRoundTripTime":0.519},"transportId":"ade48cea-76e7-4cb8-93a4-7f7141af7015","consuming":false,"producing":true},"consumingTransportStats":{"transportId":"d80b825b-015d-4b94-b9ed-79f8101cffbd","consuming":true,"producing":false},"producerStats":[{"producerId":"72b5b2e5-6e2a-4a73-8329-16e4bcfae383","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":16089,"packetsSent":1051,"nackCount":0,"remoteData":{"jitter":0,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":4,"totalRoundTripTime":0.139,"packetsLost":0}}]},{"producerId":"ffaaecfb-83ad-4815-8485-fcabd77ccb22","videoStats":[{"frameHeight":240,"frameWidth":320,"framesEncoded":504,"framesPerSecond":25,"framesSent":504,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":0,"pliCount":0,"qpSum":32065,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":21.084},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":1.795,"totalPacketSendDelay":1.795,"retransmittedBytesSent":631,"retransmittedPacketsSent":1,"bytesSent":345496,"packetsSent":518,"nackCount":1,"remoteData":{"jitter":0.000011111111111111112,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":15,"totalRoundTripTime":0.553,"packetsLost":1}},{"frameHeight":480,"frameWidth":640,"framesEncoded":504,"framesPerSecond":25,"framesSent":504,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":0,"pliCount":0,"qpSum":18260,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":21.084},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":1.777,"totalPacketSendDelay":1.777,"retransmittedBytesSent":6091,"retransmittedPacketsSent":6,"bytesSent":2170345,"packetsSent":2111,"nackCount":7,"remoteData":{"jitter":0.000011111111111111112,"fractionLost":0,"roundTripTime":0.039,"roundTripTimeMeasurements":15,"totalRoundTripTime":0.519,"packetsLost":5}}],"audioStats":[]}],"consumerStats":[]},"timestamp":"2022-07-12T05:15:55.461Z"}]},"peerId":"68ac0ae1-c101-4b8e-a08c-82028508ba8b"}`
	// test transport stats

	// test precall ends
	//precallendstr := `{"payload":{"entries":[{"event":"audio_on","timestamp":"2022-07-12T12:23:39.471Z"},{"event":"video_on","timestamp":"2022-07-12T12:23:39.471Z"},{"event":"precall_end","metaData":{"connectionInfo":{"ipDetails":{"ip":"122.161.78.105","hostname":"abts-north-dynamic-105.78.161.122.airtelbroadband.in","city":"Delhi","region":"Delhi","country":"IN","loc":"28.6519,77.2315","postal":"110001","timezone":"Asia/Kolkata","asn":{"asn":"AS24560","name":"Bharti Airtel Ltd., Telemedia Services","domain":"airtel.com","route":"122.161.76.0/22","type":"isp"},"company":{"name":"Bharti Airtel Ltd.,224, Okhla industrial Area Phase III New Delhi","domain":"airtel.com","type":"isp"},"privacy":{"vpn":false,"proxy":false,"tor":false,"relay":false,"hosting":false,"service":""},"abuse":{"address":"Bharti Airtel Ltd.","country":"IN","email":"dsltac2north.unoc@airtel.com","name":"Network Administrator for ABTS DEL","network":"122.161.0.0/16","phone":"+91-11-41615533"},"domains":{"total":0,"domains":[]}},"effectiveNetworkType":"4g","location":{"coords":{"latitude":28.6519,"longitude":77.2315}},"turnConnectivity":true,"connectivity":{"host":true,"relay":true,"reflexive":false},"throughput":4550,"fractionalLoss":0,"RTT":0.118,"jitter":0,"backendRTT":0.074}},"timestamp":"2022-07-12T12:23:41.428Z"},{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":6003,"bytesSent":886742,"packetsSent":1263,"packetsReceived":84,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.038,"totalRoundTripTime":0.038},"transportId":"973a7ead-20a3-4900-8748-6828f57cde5e","consuming":false,"producing":true},"consumingTransportStats":{"transportId":"7f697819-6795-4f91-8b42-4d4e6ac4af42","consuming":true,"producing":false},"producerStats":[null,{"producerId":"0b15f59e-4c48-4a1f-abfe-7ff7ddf30e59","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":11868,"packetsSent":365,"nackCount":0,"remoteData":{"jitter":0,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":1,"totalRoundTripTime":0.038,"packetsLost":0}}]}],"consumerStats":[]},"timestamp":"2022-07-12T12:23:46.790Z"}]},"peerId":"a60b7be3-3ee7-4229-ae48-b5251a11ce1a"}`
	//filteredValues := filterSliceForNonNull(processEntries(gjson.Get(precallendstr, "peerId").String(), gjson.Get(precallendstr, "payload")))
	//fmt.Println(filteredValues)
	// test peer information
	//peerinformationstr := `{"payload":{"entries":[{"event":"precall_begin","timestamp":"2022-07-12T14:17:48.032Z"},{"event":"tab_change","backToCurrentTab":true,"timestamp":"2022-07-12T14:17:48.255Z"},{"event":"browser_foregrounded","timestamp":"2022-07-12T14:17:48.255Z"},{"event":"call_join","metaData":{"peerMetaData":{"userId":"7efb4219-97f7-4a1d-914e-1470ce2bcf3f","peerId":"12d4affb-0ec6-4212-8f37-ae08886e169a","displayName":" chikmo","roomViewType":"groupCall","roomName":"uwnbea-humkxm","deviceInfo":{"isMobile":false,"browserName":"Chrome","osName":"macOS","browserVersion":"103.0.0.0","osVersionName":"10.15.7","engineName":"Blink","userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36","memory":8,"cpus":10},"metaData":{},"permissions":{},"meetingEnv":"production"}},"timestamp":"2022-07-12T14:17:54.682Z"}]},"peerId":"12d4affb-0ec6-4212-8f37-ae08886e169a"}`
	//filteredValues := filterSliceForNonNull(processEntries(gjson.Get(peerinformationstr, "peerId").String(), gjson.Get(peerinformationstr, "payload")))
	//fmt.Println(filteredValues)

	// test producer audio stats
	//producerstatsstr := `{"payload":{"entries":[{"event":"audio_on","timestamp":"2022-07-12T12:23:39.471Z"},{"event":"video_on","timestamp":"2022-07-12T12:23:39.471Z"},{"event":"precall_end","metaData":{"connectionInfo":{"ipDetails":{"ip":"122.161.78.105","hostname":"abts-north-dynamic-105.78.161.122.airtelbroadband.in","city":"Delhi","region":"Delhi","country":"IN","loc":"28.6519,77.2315","postal":"110001","timezone":"Asia/Kolkata","asn":{"asn":"AS24560","name":"Bharti Airtel Ltd., Telemedia Services","domain":"airtel.com","route":"122.161.76.0/22","type":"isp"},"company":{"name":"Bharti Airtel Ltd.,224, Okhla industrial Area Phase III New Delhi","domain":"airtel.com","type":"isp"},"privacy":{"vpn":false,"proxy":false,"tor":false,"relay":false,"hosting":false,"service":""},"abuse":{"address":"Bharti Airtel Ltd.","country":"IN","email":"dsltac2north.unoc@airtel.com","name":"Network Administrator for ABTS DEL","network":"122.161.0.0/16","phone":"+91-11-41615533"},"domains":{"total":0,"domains":[]}},"effectiveNetworkType":"4g","location":{"coords":{"latitude":28.6519,"longitude":77.2315}},"turnConnectivity":true,"connectivity":{"host":true,"relay":true,"reflexive":false},"throughput":4550,"fractionalLoss":0,"RTT":0.118,"jitter":0,"backendRTT":0.074}},"timestamp":"2022-07-12T12:23:41.428Z"},{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":6003,"bytesSent":886742,"packetsSent":1263,"packetsReceived":84,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.038,"totalRoundTripTime":0.038},"transportId":"973a7ead-20a3-4900-8748-6828f57cde5e","consuming":false,"producing":true},"consumingTransportStats":{"transportId":"7f697819-6795-4f91-8b42-4d4e6ac4af42","consuming":true,"producing":false},"producerStats":[null,{"producerId":"0b15f59e-4c48-4a1f-abfe-7ff7ddf30e59","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":11868,"packetsSent":365,"nackCount":0,"remoteData":{"jitter":0,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":1,"totalRoundTripTime":0.038,"packetsLost":0}}]}],"consumerStats":[]},"timestamp":"2022-07-12T12:23:46.790Z"}]},"peerId":"a60b7be3-3ee7-4229-ae48-b5251a11ce1a"}`
	//filteredValues := filterSliceForNonNull(processEntries(gjson.Get(producerstatsstr, "peerId").String(), gjson.Get(producerstatsstr, "payload")))
	//fmt.Println(filteredValues)

	// test producer video stats
	//videostatstr := `{"payload":{"entries":[{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":53060,"bytesSent":8056292,"packetsSent":11147,"packetsReceived":846,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.038,"totalRoundTripTime":0.492},"transportId":"4a7e5e27-1ba1-41b4-a386-7b6ed5f116fe","consuming":false,"producing":true},"consumingTransportStats":{"transportId":"a375600d-e910-4c17-a03e-33b93fd810a2","consuming":true,"producing":false},"producerStats":[{"producerId":"a6d0931d-efff-4b77-aacc-6264af4aa9f1","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":113509,"packetsSent":3149,"nackCount":0,"remoteData":{"jitter":0,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":12,"totalRoundTripTime":0.492,"packetsLost":2}}]},{"producerId":"c092ffa0-5ad1-4074-95db-84b35f8363b3","videoStats":[{"frameHeight":240,"frameWidth":320,"framesEncoded":1509,"framesPerSecond":24,"framesSent":1509,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":1,"pliCount":0,"qpSum":71294,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":63.045},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":6.795,"totalPacketSendDelay":6.795,"retransmittedBytesSent":4212,"retransmittedPacketsSent":6,"bytesSent":1052863,"packetsSent":1551,"nackCount":30,"remoteData":{"jitter":0.000044444444444444447,"fractionLost":0,"roundTripTime":0.038,"roundTripTimeMeasurements":49,"totalRoundTripTime":2.101,"packetsLost":8}},{"frameHeight":480,"frameWidth":640,"framesEncoded":1509,"framesPerSecond":24,"framesSent":1509,"keyFramesEncoded":1,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":0,"pliCount":0,"qpSum":34681,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":63.045},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":6.75,"totalPacketSendDelay":6.75,"retransmittedBytesSent":58894,"retransmittedPacketsSent":56,"bytesSent":6420855,"packetsSent":6232,"nackCount":94,"remoteData":{"jitter":0.000033333333333333335,"fractionLost":0,"roundTripTime":0.037,"roundTripTimeMeasurements":49,"totalRoundTripTime":2.153,"packetsLost":51}}],"audioStats":[]}],"consumerStats":[]},"timestamp":"2022-07-11T05:00:22.701Z"}]},"peerId":"8e94afa8-fcbe-4ed3-b3ce-1e040fec477e"}`
	//filteredValues := filterSliceForNonNull(processEntries(gjson.Get(videostatstr, "peerId").String(), gjson.Get(videostatstr, "payload")))
	//fmt.Println(filteredValues)

	// test consumer audio stats
	consumerStatsStr := `{"payload":{"entries":[{"event":"ping_stat","metaData":{"producingTransportStats":{"stats":{"bytesReceived":227133,"bytesSent":39008666,"packetsSent":53687,"packetsReceived":3490,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0.033,"totalRoundTripTime":8.312},"transportId":"ccb42ea1-366d-4af5-bbae-3865c1be17cd","consuming":false,"producing":true},"consumingTransportStats":{"stats":{"bytesReceived":6834459,"bytesSent":75376,"packetsSent":1628,"packetsReceived":6700,"dtlsCipher":"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256","dtlsState":"connected","iceRole":"controlling","roundTripTime":0,"totalRoundTripTime":0},"transportId":"fa5fe644-6365-4dd6-bb6b-2aca27244a87","consuming":true,"producing":false},"producerStats":[{"producerId":"483dedaf-bbf4-4e54-bec9-fa05d5d23703","videoStats":[],"audioStats":[{"retransmittedBytesSent":0,"retransmittedPacketsSent":0,"bytesSent":533144,"packetsSent":15049,"nackCount":0,"remoteData":{"jitter":0.00008333333333333333,"fractionLost":0,"roundTripTime":0.03,"roundTripTimeMeasurements":61,"totalRoundTripTime":2.614,"packetsLost":32}}]},{"producerId":"aae42bbf-664e-4f5d-b281-ff13ccb424fc","videoStats":[{"frameHeight":480,"frameWidth":640,"framesEncoded":7216,"framesPerSecond":24,"framesSent":7216,"keyFramesEncoded":10,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":3,"pliCount":3,"qpSum":112110,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":301.04},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":27.652,"totalPacketSendDelay":27.652,"retransmittedBytesSent":45798,"retransmittedPacketsSent":43,"bytesSent":31122422,"packetsSent":30193,"nackCount":22,"remoteData":{"jitter":0.000011111111111111112,"fractionLost":0,"roundTripTime":0.031,"roundTripTimeMeasurements":221,"totalRoundTripTime":8.17,"packetsLost":56}},{"frameHeight":240,"frameWidth":320,"framesEncoded":7216,"framesPerSecond":24,"framesSent":7216,"keyFramesEncoded":10,"firCount":0,"encoderImplementation":"libvpx","hugeFramesSent":2,"pliCount":5,"qpSum":228173,"qualityLimitationDurations":{"other":0,"cpu":0,"bandwidth":0,"none":301.04},"qualityLimitationReason":"none","qualityLimitationResolutionChanges":0,"totalEncodeTime":27.881,"totalPacketSendDelay":27.881,"retransmittedBytesSent":6168,"retransmittedPacketsSent":9,"bytesSent":5096179,"packetsSent":7393,"nackCount":30,"remoteData":{"jitter":0.00006666666666666667,"fractionLost":0,"roundTripTime":0.033,"roundTripTimeMeasurements":221,"totalRoundTripTime":8.312,"packetsLost":14}}],"audioStats":[]}],"consumerStats":[{"consumerId":"ceabf113-dee1-49b1-82c5-29188027049b","peerId":"6b9ad0fe-1346-462f-beef-ed230bac7629","producerId":"78f32abb-9559-4e2c-a95f-3b1deddc1efa","videoStats":[],"audioStats":[{"audioLevel":0,"concealmentEvents":5,"jitterBufferDelay":15062.4,"jitterBufferEmittedCount":258240,"totalAudioEnergy":2.55644792006764,"totalSamplesDuration":5.339999999999931,"totalSamplesReceived":256320,"bytesReceived":11421,"packetsReceived":273,"packetsLost":0,"jitter":0.01}]},{"consumerId":"ef1c7d9f-43b3-4e10-9a2c-02e6173d09f7","peerId":"6b9ad0fe-1346-462f-beef-ed230bac7629","producerId":"f3e75c80-f851-4001-a0dd-3c7c53c2a173","videoStats":[{"frameHeight":640,"frameWidth":480,"framesDecoded":1535,"framesDropped":1,"framesPerSecond":23,"framesReceived":1536,"keyFramesDecoded":6,"firCount":0,"decoderImplementation":"libvpx","bytesReceived":6462853,"packetsReceived":6249,"packetsLost":0,"jitter":0.024,"nackCount":0}],"audioStats":[]}]},"timestamp":"2022-07-12T16:27:17.467Z"}]},"peerId":"cbc701ba-583a-46a5-a9c4-65127b439e12"}`
	filteredValues := filterSliceForNonNull(processEntries(gjson.Get(consumerStatsStr, "peerId").String(), gjson.Get(consumerStatsStr, "payload")))
	fmt.Println(filteredValues)

}
