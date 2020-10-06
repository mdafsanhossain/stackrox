package util

import io.stackrox.proto.storage.NetworkFlowOuterClass
import io.stackrox.proto.storage.NetworkFlowOuterClass.NetworkEntityInfo
import objects.Edge
import io.stackrox.proto.api.v1.NetworkGraphOuterClass

class NetworkGraphUtil {

    static int edgeCount(NetworkGraphOuterClass.NetworkGraph graph) {
        int numEdges = 0
        graph.nodesList.each {
            numEdges += it.outEdgesCount
        }
        return numEdges
    }

    static Set<String> deployments(NetworkGraphOuterClass.NetworkGraph graph) {
        def deploymentSet = new HashSet<String>([])

        graph.nodesList.each {
            if (it.entity.type != NetworkFlowOuterClass.NetworkEntityInfo.Type.DEPLOYMENT) {
                return
            }
            deploymentSet.add("${it.entity.deployment.namespace}/${it.entity.deployment.name}")
        }
        return deploymentSet
    }

    private static String entityLabel(NetworkEntityInfo entity) {
        if (entity.type == NetworkFlowOuterClass.NetworkEntityInfo.Type.DEPLOYMENT) {
            return "${entity.deployment.namespace}/${entity.deployment.name}"
        } else if (entity.type == NetworkFlowOuterClass.NetworkEntityInfo.Type.INTERNET) {
            return "INTERNET"
        }
        return ""
    }

    static Set<String> flowStrings(NetworkGraphOuterClass.NetworkGraph graph) {
        return new HashSet<String>(graph.nodesList.<String>collectMany {
            def srcLabel = entityLabel(it.entity)
            return srcLabel ? it.outEdges.collectMany {
                def tgt = graph.nodesList.get(it.key)
                def dstLabel = entityLabel(tgt.entity)
                return dstLabel ? ["${srcLabel} -> ${dstLabel}"] : []
            } : []
        })
    }

    static NetworkGraphOuterClass.NetworkNode findDeploymentNode(
            NetworkGraphOuterClass.NetworkGraph graph, String deploymentId) {
        return graph.nodesList.find {
            it.deploymentId == deploymentId
        }
    }

    static List<Edge> findEdges(NetworkGraphOuterClass.NetworkGraph graph, String sourceId, String targetId) {
        println "Checking for edge between deployments: sourceId ${sourceId}, targetId ${targetId}"

        def sourceNodes = sourceId == null ? graph.nodesList : graph.nodesList.findAll {
            it.deploymentId == sourceId
        }
        def targetNodeIndex = graph.nodesList.findIndexOf {
            it.deploymentId == targetId
        }

        if ((sourceId != null && sourceNodes.empty) || (targetId != null && targetNodeIndex == -1)) {
            if (sourceId != null && sourceNodes.empty) {
                println "Found no nodes matching sourceId ${sourceId}"
            }
            if (targetId != null && targetNodeIndex == -1) {
                println "Found no nodes matching targetId ${targetId}"
            }
            return []
        }

        println "Looking at edges for ${sourceNodes.size()} source node(s)"

        return sourceNodes.collectMany {
            def currentSourceId = it.deploymentId
            return it.getOutEdgesMap().collectMany {
                if (targetNodeIndex != -1 && it.key != targetNodeIndex) {
                    return []
                }
                println "Source Id ${currentSourceId} -> edge target key: ${it.key}"
                def targetNode = graph.nodesList.get(it.key)
                println "  -> targetId: ${targetNode.deploymentId}"

                def props = it.value.propertiesList
                props.forEach {
                    edgeProp -> println "    -> edge: ${edgeProp.port} ${edgeProp.protocol} "+
                            "${edgeProp.lastActiveTimestamp.seconds}.${edgeProp.lastActiveTimestamp.nanos}"
                }
                if (props == null || props.empty) {
                    props = [null]
                }
                props.collect {
                    new Edge(sourceID: currentSourceId, targetID: targetNode.deploymentId, edgeProperties: it)
                }
            }
        }
    }

}

