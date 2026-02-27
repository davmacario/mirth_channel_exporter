package exporter

/*
*
*
* <map>
*   <entry>
*     <string>date</string>
*     <string>May 28, 2025</string>
*   </entry>
*   <entry>
*     <string>channelCount</string>
*     <int>4</int>
*   </entry>
*   <entry>
*     <string>database</string>
*     <string>postgres</string>
*   </entry>
*   <entry>
*     <string>connectors</string>
*     <map>
*       <entry>
*         <string>SMTP Sender</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>File Writer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>HTTP Sender</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>File Reader</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>DICOM Listener</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>HTTP Listener</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Channel Writer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JavaScript Reader</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>TCP Listener</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Database Reader</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Database Writer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Document Writer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Web Service Listener</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JavaScript Writer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>TCP Sender</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Channel Reader</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Web Service Sender</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JMS Sender</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JMS Listener</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>DICOM Sender</string>
*         <string>4.6.1</string>
*       </entry>
*     </map>
*   </entry>
*   <entry>
*     <string>plugins</string>
*     <map>
*       <entry>
*         <string>Server Log</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Text Viewer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Delimited Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>XML Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>HTTP Authentication Settings</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Channel History</string>
*         <string>4.6.0</string>
*       </entry>
*       <entry>
*         <string>Image Viewer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>HL7v2 Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Data Pruner</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Raw Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Message Generator</string>
*         <string>4.6.0</string>
*       </entry>
*       <entry>
*         <string>DICOM Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Destination Set Filter Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>HL7v3 Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Command Center Base</string>
*         <string>4.6.0</string>
*       </entry>
*       <entry>
*         <string>Command Center Analytics</string>
*         <string>4.6.0</string>
*       </entry>
*       <entry>
*         <string>JavaScript Filter Rule</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Mapper Transformer Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Message Builder Transformer Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JavaScript Transformer Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Dashboard Connector Status Monitor</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>DICOM Viewer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Directory Resource Plugin</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>External Script Transformer Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>External Script Filter Step</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>NCPDP Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Transmission Mode - MLLP</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Rule Builder Filter Rule</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>TCP Connector Service Plugin</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>JSON Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>License Manager Plugin</string>
*         <string>4.6.0</string>
*       </entry>
*       <entry>
*         <string>PDF Viewer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>EDI Data Type</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>SSL Manager</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>Global Map Viewer</string>
*         <string>4.6.1</string>
*       </entry>
*       <entry>
*         <string>XSLT Transformer Step</string>
*         <string>4.6.1</string>
*       </entry>
*     </map>
*   </entry>
*   <entry>
*     <string>name</string>
*     <null/>
*   </entry>
*   <entry>
*     <string>version</string>
*     <string>4.6.1</string>
*   </entry>
* </map>%
 */
type ServerStats struct {
}
