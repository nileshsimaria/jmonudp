# jmonudp
Junos UDP Telemetry client

# Sample data dump in json for qmon sensor.
<pre>
$ ./jmonudp --help
Usage of /tmp/jmonudp:
      --port int   UDP port number to listen on
      
$ ./jmonudp --port 27000 
{           
  "queue_monitor_element_info": [
    {       
      "if_name": "xe-2/0/1",
      "queue_monitor_stats_ingress": {
        "queue_monitor_stats_info": [
          {
            "queue_number": 0,
            "queue_id": 32,
            "peak_buffer_occupancy_bytes": 1321,
            "peak_buffer_occupancy_percent": 0,
            "packets": 55526232,
            "octets": 73294626240,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },
          {
            "queue_number": 1,
            "queue_id": 33,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },
          { 
            "queue_number": 2,
            "queue_id": 34,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },
          { 
            "queue_number": 3,
            "queue_id": 35,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 98,
            "octets": 8232,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 4,
            "queue_id": 36,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 5,
            "queue_id": 37,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 6,
            "queue_id": 38,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 7,
            "queue_id": 39,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
            "red_drop_octets_color_0": 0,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          }    
        ]    
      },   
      "queue_monitor_stats_egress": {
        "queue_monitor_stats_info": [
          {    
            "queue_number": 0,
            "queue_id": 24,
            "peak_buffer_occupancy_bytes": 61867872,
            "peak_buffer_occupancy_percent": 99,
            "packets": 27166602467,
            "octets": 35859915137508,
            "tail_drop_packets": 729, 
            "tail_drop_octets": 962280,
            "red_drop_packets_color_0": 27571518,
            "red_drop_octets_color_0": 36394403760,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 1,
            "queue_id": 25,
            "peak_buffer_occupancy_bytes": 56147586,
            "peak_buffer_occupancy_percent": 98,
            "packets": 75966079936,
            "octets": 100275225516840,
            "tail_drop_packets": 451, 
            "tail_drop_octets": 595320,
            "red_drop_packets_color_0": 22151825,
            "red_drop_octets_color_0": 29240409000,
            "red_drop_packets_color_1": 0,
            "red_drop_octets_color_1": 0,
            "red_drop_packets_color_2": 0,
            "red_drop_octets_color_2": 0,
            "red_drop_packets_color_3": 0,
            "red_drop_octets_color_3": 0
          },   
          {    
            "queue_number": 2,
            "queue_id": 26,
            "peak_buffer_occupancy_bytes": 0,
            "peak_buffer_occupancy_percent": 0,
            "packets": 0,
            "octets": 0,
            "tail_drop_packets": 0,
            "tail_drop_octets": 0,
            "red_drop_packets_color_0": 0,
...
...
...
</pre>
