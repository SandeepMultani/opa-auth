package auth

import future.keywords

# rule to return the resource instances by ids
tenants[id] := tenant_instance {
	some tenant_instance in data.tenants
    id := sprintf("tenant:%s",[tenant_instance.id])
}

domains[id] := domain_instance {
	some domain_instance in data.domains
	id := sprintf("domain:%s",[domain_instance.id])
}

entities[id] := entity_instance {
	some entity_instance in data.entities
	id := sprintf("domain_entity:%s",[entity_instance.id])
}

# return a full graph mapping of each subject to the object it has reference to
full_graph[subject] := ref_object {
	some subject, object_instance in object.union_n([tenants, domains, entities])

	# get the parent_id the subject is referring
	ref_object := [object.get(object_instance, "parent_id", null)]
}

# rule to return users by ids
users[id] := user {
	some user in data.users
	id := user.id
}

# the input user
input_user := users[input.user]

# rule to return a list of allowed assignments
allowing_assignments[assignment] {
	# iterate the user assignments
	some assignment in input_user.assignments

	# check that the required action from the input is allowed by the current role
	input.action in data.roles[assignment.role].grants

	# check that the required resource from the input is reachable in the graph
	# by the current team 
	assignment.resource in graph.reachable(full_graph, {input.resource})
}

# create allow rule with the default of false
default allow := false

allow {
	# allow the user to perform the action on the resource if they have more than one allowing assignments
	count(allowing_assignments) > 0
}