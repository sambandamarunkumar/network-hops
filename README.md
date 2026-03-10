# network-hops
**Data Locality Optimization for Low Latency Distributed Systems**

### Paper Information
- **Author(s):** Arunkumar Sambandam
- **Published In:** Journal of Advances in Developmental Research(IJAIDR)
- **Publication Date:** Sept 1, 2021
- **ISSN:** E-ISSN: 0976-4844
- **DOI:**
- **Impact Factor:**9.71

### Abstract
Distributed systems achieve scalability through partitioned data placement across multiple nodes, but as clusters grow, requests often traverse multiple intermediate machines. This multi‑hop communication introduces routing overhead, longer paths, and increased network delay. Each additional hop adds propagation time, congestion, and reduces efficiency, while static placement fails to adapt to evolving access patterns. The accumulation of unnecessary traversals increases latency variability and limits predictable performance. Minimizing hop distance between clients and data is therefore critical for improving communication efficiency and scalability in distributed architectures.

## Key Contributions
- **Hop Count Analysis:**
   Quantified how multi‑hop communication increases routing overhead, latency, and congestion in distributed systems.

- **Locality‑Aware Placement:**
   Proposed adaptive partitioning strategies that prioritize proximity, reducing unnecessary inter‑node traversal.

- **Performance Modeling:**
   Demonstrated that average hop count rises with cluster size, limiting scalability despite added resources.

- **Simulation & Validation:**
   Implemented prototype models showing reduced hop counts and improved response times compared to static placement.

- **Communication Efficiency Framework:**
   Established a structured methodology for minimizing traversal distance and stabilizing latency variability.

## Relevance & Impact
- **Lower Latency:**
   Reduced hop counts shorten transmission paths, cutting propagation delays and switch processing overhead.

- **Improved Throughput:**
   Balanced communication paths sustain higher request completion rates under heavy workloads.

- **Scalable Performance:**
   Maintains efficiency as clusters expand, avoiding communication‑dominated bottlenecks.

- **Cost & Energy Efficiency:**
   Less inter‑node traffic reduces bandwidth usage, energy consumption, and infrastructure overhead.

- **Predictable Behavior:**
   Minimizing traversal distance stabilizes latency variability, ensuring consistent service quality.

### Experimental Results (Summary)

  | Nodes | Static Placement Hops | Locality aware placement hops | Improvment (%)|
  |-------|-----------------------| ------------------------------| --------------|
  | 3     |  3.2                  | 1.7                           | 47            |
  | 5     |  4.0                  | 2.0                           | 50            |
  | 7     |  4.8                  | 2.3                           | 52            |
  | 9     |  5.6                  | 2.6                           | 54            |
  | 11    |  6.3                  | 2.9                           | 54            |

### Citation
Data Locality Optimization for Low Latency Distributed Systems
* Arunkumar Sambandam
* Journal of Advances in Developmental Research(IJAIDR)
* ISSN E-ISSN: E-ISSN: 0976-4844
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
[https://www.ijlrp.com*****************/ \](https://www.ijaidr.com/research-paper.php?id=1707) \
**Author Contact** \
**LinkedIn**: linkedin.com/in/arunkumar-sambandam-9b769b6 | **Email**: arunkumar.sambandam@yahoo.com







