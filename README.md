# network-hops
**Data Locality Optimization for Low Latency Distributed Systems**

### Paper Information
- **Author(s):** Arunkumar Sambandam
- **Published In:** *********************************************International Journal of Leading Research Publication (IJLRP)
- **Publication Date:** ******Aug 2021
- **ISSN:** E-ISSN: **********2582-8010
- **DOI:**
- **Impact Factor:** *******9.56

### Abstract
This paper addresses the challenge of identifying Input/Output bottlenecks in complex distributed data pipelines where isolated observability metrics provide limited insight. 
It highlights how single-modal monitoring fails to accurately localize performance issues caused by cross-layer interactions in dynamic, heterogeneous environments. The proposed 
Multimodal Observability framework unifies metrics, logs, and traces to enable correlated, cross-layer analysis of Input/Output behavior. By aligning low-level signals with 
execution paths, the approach improves bottleneck detection accuracy and reduces misdiagnosis in large-scale pipelines.

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

  | Nodes | Baseline Request Completion Time (ms) | Multimodal Request Completion Tim (ms) | Improvment (%)  |
  |-------|---------------------------------------| ---------------------------------------| ----------------|
  | 3     |  120                                  | 95                                     | 20.83           |
  | 5     |  145                                  | 115                                    | 20.69           |
  | 7     |  175                                  | 140                                    | 20.00           |
  | 9     |  210                                  | 165                                    | 21.43           |
  | 11    |  250                                  | 195                                    | 22.00           |

### Citation
Multimodal Observability for Input Output Bottleneck Detection
* Arunkumar Sambandam
* ***********************************International Journal of Leading Research Publication 
* ISSN E-ISSN: *****************************2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
https://www.ijlrp.com*****************/ \
**Author Contact** \
**LinkedIn**: https://www.linkedin.com/**** | **Email**: arunkumar.sambandam@yahoo.com







